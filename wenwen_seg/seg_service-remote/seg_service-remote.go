// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
	"wenwen_seg"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  QuerySegResponse query_segment(string query_gbk, i32 query_info)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_type(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_wtype(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_wf(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_qf(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_qef(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  i32 get_term_delf(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  string get_term_text(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  string get_term_text_gbk(i32 term_id)")
	fmt.Fprintln(os.Stderr, "  string postag2string(i32 wtype)")
	fmt.Fprintln(os.Stderr, "   get_extend_term(i32 term_id)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := wenwen_seg.NewSegServiceClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "query_segment":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "QuerySegment requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		tmp1, err36 := (strconv.Atoi(flag.Arg(2)))
		if err36 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.QuerySegment(value0, value1))
		fmt.Print("\n")
		break
	case "get_term_type":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermType requires 1 args")
			flag.Usage()
		}
		tmp0, err37 := (strconv.Atoi(flag.Arg(1)))
		if err37 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermType(value0))
		fmt.Print("\n")
		break
	case "get_term_wtype":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermWtype requires 1 args")
			flag.Usage()
		}
		tmp0, err38 := (strconv.Atoi(flag.Arg(1)))
		if err38 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermWtype(value0))
		fmt.Print("\n")
		break
	case "get_term_wf":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermWf requires 1 args")
			flag.Usage()
		}
		tmp0, err39 := (strconv.Atoi(flag.Arg(1)))
		if err39 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermWf(value0))
		fmt.Print("\n")
		break
	case "get_term_qf":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermQf requires 1 args")
			flag.Usage()
		}
		tmp0, err40 := (strconv.Atoi(flag.Arg(1)))
		if err40 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermQf(value0))
		fmt.Print("\n")
		break
	case "get_term_qef":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermQef requires 1 args")
			flag.Usage()
		}
		tmp0, err41 := (strconv.Atoi(flag.Arg(1)))
		if err41 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermQef(value0))
		fmt.Print("\n")
		break
	case "get_term_delf":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermDelf requires 1 args")
			flag.Usage()
		}
		tmp0, err42 := (strconv.Atoi(flag.Arg(1)))
		if err42 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermDelf(value0))
		fmt.Print("\n")
		break
	case "get_term_text":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermText requires 1 args")
			flag.Usage()
		}
		tmp0, err43 := (strconv.Atoi(flag.Arg(1)))
		if err43 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermText(value0))
		fmt.Print("\n")
		break
	case "get_term_text_gbk":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTermTextGbk requires 1 args")
			flag.Usage()
		}
		tmp0, err44 := (strconv.Atoi(flag.Arg(1)))
		if err44 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetTermTextGbk(value0))
		fmt.Print("\n")
		break
	case "postag2string":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Postag2string requires 1 args")
			flag.Usage()
		}
		tmp0, err45 := (strconv.Atoi(flag.Arg(1)))
		if err45 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.Postag2string(value0))
		fmt.Print("\n")
		break
	case "get_extend_term":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetExtendTerm requires 1 args")
			flag.Usage()
		}
		tmp0, err46 := (strconv.Atoi(flag.Arg(1)))
		if err46 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetExtendTerm(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
