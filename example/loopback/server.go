/*
 * Copyright 2022 Han Xin, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/Mystarset/demo/fuse2grpc"
	"github.com/Mystarset/demo/pb"
	"github.com/Mystarset/demo/pkg/utils"
)

func main() {
	debug := flag.Bool("debug", false, "print debugging messages.")
	other := flag.Bool("allow-other", false, "mount with -o allowother.")
	quiet := flag.Bool("q", false, "quiet")
	ro := flag.Bool("ro", false, "mount read-only")
	loggerLevel := flag.String("logger-level", "info", "log level")
	flag.Parse()

	if flag.NArg() < 1 {
		logrus.Fatalf("Usage: %s <ORIGINAL>", path.Base(os.Args[0]))
	}

	logrus.SetLevel(utils.GetLogLevel(*loggerLevel))
	orig := flag.Arg(0)

	l, err := net.Listen("tcp", "127.0.0.1:8760")
	if err != nil {
		logrus.Fatal(err)
	}

	logEntry := logrus.NewEntry(logrus.StandardLogger())
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logEntry),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logEntry),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	grpc_prometheus.Register(s)

	loopbackRoot, err := fs.NewLoopbackRoot(orig)
	if err != nil {
		logrus.Fatalf("NewLoopbackRoot: %v", err)
	}

	sec := time.Second
	opts := &fs.Options{
		// These options are to be compatible with libfuse defaults,
		// making benchmarking easier.
		AttrTimeout:  &sec,
		EntryTimeout: &sec,
	}
	opts.Debug = *debug
	opts.AllowOther = *other
	if opts.AllowOther {
		// Make the kernel check file permissions for us
		opts.MountOptions.Options = append(opts.MountOptions.Options, "default_permissions")
	}
	if *ro {
		opts.MountOptions.Options = append(opts.MountOptions.Options, "ro")
	}
	// First column in "df -T": original dir
	opts.MountOptions.Options = append(opts.MountOptions.Options, "fsname="+orig)
	// Second column in "df -T" will be shown as "fuse." + Name
	opts.MountOptions.Name = "loopback"
	// Leave file permissions on "000" files as-is
	opts.NullPermissions = true
	// Enable diagnostics logging
	if !*quiet {
		opts.Logger = log.New(os.Stderr, "", 0)
	}

	rawFS := fs.NewNodeFS(loopbackRoot, opts)

	srv := fuse2grpc.NewServer(rawFS)

	pb.RegisterRawFileSystemServer(s, srv)
	go s.Serve(l)

	logrus.Infof("Listen on %s for dir %s", l.Addr(), orig)

	signal.Ignore(syscall.SIGPIPE)
	sigCh := make(chan os.Signal, 10)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	for range sigCh {
		s.Stop()
		logrus.Info("Shutdon")
		return
	}
}
