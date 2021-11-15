// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
	tts "github.com/cloudwego/kitex/output/kitex_gen/thrifttest/thrifttest"
)

func main() {
	cli := tts.MustNewClient("a.b.c", client.WithHostPorts("[::1]:8888"))
	fmt.Println(cli.TestMap(context.Background(), map[int32]int32{123: 123}))
	fmt.Println(cli.TestStringMap(context.Background(), map[string]string{"123": "123"}))
	fmt.Println(cli.TestMapMap(context.Background(), 555))
}
