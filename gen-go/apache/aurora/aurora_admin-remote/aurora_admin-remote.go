// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"apache/aurora"
	"context"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  Response setQuota(string ownerRole, ResourceAggregate quota)")
	fmt.Fprintln(os.Stderr, "  Response forceTaskState(string taskId, ScheduleStatus status)")
	fmt.Fprintln(os.Stderr, "  Response performBackup()")
	fmt.Fprintln(os.Stderr, "  Response listBackups()")
	fmt.Fprintln(os.Stderr, "  Response stageRecovery(string backupId)")
	fmt.Fprintln(os.Stderr, "  Response queryRecovery(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response deleteRecoveryTasks(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response commitRecovery()")
	fmt.Fprintln(os.Stderr, "  Response unloadRecovery()")
	fmt.Fprintln(os.Stderr, "  Response startMaintenance(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response drainHosts(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response maintenanceStatus(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response endMaintenance(Hosts hosts)")
	fmt.Fprintln(os.Stderr, "  Response slaDrainHosts(Hosts hosts, SlaPolicy defaultSlaPolicy, i64 timeoutSecs)")
	fmt.Fprintln(os.Stderr, "  Response snapshot()")
	fmt.Fprintln(os.Stderr, "  Response triggerExplicitTaskReconciliation(ExplicitReconciliationSettings settings)")
	fmt.Fprintln(os.Stderr, "  Response triggerImplicitTaskReconciliation()")
	fmt.Fprintln(os.Stderr, "  Response pruneTasks(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response createJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response scheduleCronJob(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response descheduleCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response startCronJob(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response restartShards(JobKey job,  shardIds)")
	fmt.Fprintln(os.Stderr, "  Response killTasks(JobKey job,  instances, string message)")
	fmt.Fprintln(os.Stderr, "  Response addInstances(InstanceKey key, i32 count)")
	fmt.Fprintln(os.Stderr, "  Response replaceCronTemplate(JobConfiguration config)")
	fmt.Fprintln(os.Stderr, "  Response startJobUpdate(JobUpdateRequest request, string message)")
	fmt.Fprintln(os.Stderr, "  Response pauseJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response resumeJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response abortJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response rollbackJobUpdate(JobUpdateKey key, string message)")
	fmt.Fprintln(os.Stderr, "  Response pulseJobUpdate(JobUpdateKey key)")
	fmt.Fprintln(os.Stderr, "  Response getRoleSummary()")
	fmt.Fprintln(os.Stderr, "  Response getJobSummary(string role)")
	fmt.Fprintln(os.Stderr, "  Response getTasksStatus(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getTasksWithoutConfigs(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getPendingReason(TaskQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getConfigSummary(JobKey job)")
	fmt.Fprintln(os.Stderr, "  Response getJobs(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response getQuota(string ownerRole)")
	fmt.Fprintln(os.Stderr, "  Response populateJobConfig(JobConfiguration description)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateSummaries(JobUpdateQuery jobUpdateQuery)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDetails(JobUpdateQuery query)")
	fmt.Fprintln(os.Stderr, "  Response getJobUpdateDiff(JobUpdateRequest request)")
	fmt.Fprintln(os.Stderr, "  Response getTierConfigs()")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func (h httpHeaders) Set(value string) error {
	parts := strings.Split(value, ": ")
	if len(parts) != 2 {
		return fmt.Errorf("header should be of format 'Key: Value'")
	}
	h[parts[0]] = parts[1]
	return nil
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	headers := make(httpHeaders)
	var parsedUrl *url.URL
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
	flag.Var(headers, "H", "Headers to set on the http(s) request (e.g. -H \"Key: Value\")")
	flag.Parse()

	if len(urlString) > 0 {
		var err error
		parsedUrl, err = url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http" || parsedUrl.Scheme == "https"
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
		if len(headers) > 0 {
			httptrans := trans.(*thrift.THttpClient)
			for key, value := range headers {
				httptrans.SetHeader(key, value)
			}
		}
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
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := aurora.NewAuroraAdminClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "setQuota":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetQuota requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		arg355 := flag.Arg(2)
		mbTrans356 := thrift.NewTMemoryBufferLen(len(arg355))
		defer mbTrans356.Close()
		_, err357 := mbTrans356.WriteString(arg355)
		if err357 != nil {
			Usage()
			return
		}
		factory358 := thrift.NewTJSONProtocolFactory()
		jsProt359 := factory358.GetProtocol(mbTrans356)
		argvalue1 := aurora.NewResourceAggregate()
		err360 := argvalue1.Read(jsProt359)
		if err360 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		fmt.Print(client.SetQuota(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "forceTaskState":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ForceTaskState requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		tmp1, err := (strconv.Atoi(flag.Arg(2)))
		if err != nil {
			Usage()
			return
		}
		argvalue1 := aurora.ScheduleStatus(tmp1)
		value1 := argvalue1
		fmt.Print(client.ForceTaskState(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "performBackup":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "PerformBackup requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.PerformBackup(context.Background()))
		fmt.Print("\n")
		break
	case "listBackups":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "ListBackups requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.ListBackups(context.Background()))
		fmt.Print("\n")
		break
	case "stageRecovery":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StageRecovery requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.StageRecovery(context.Background(), value0))
		fmt.Print("\n")
		break
	case "queryRecovery":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "QueryRecovery requires 1 args")
			flag.Usage()
		}
		arg363 := flag.Arg(1)
		mbTrans364 := thrift.NewTMemoryBufferLen(len(arg363))
		defer mbTrans364.Close()
		_, err365 := mbTrans364.WriteString(arg363)
		if err365 != nil {
			Usage()
			return
		}
		factory366 := thrift.NewTJSONProtocolFactory()
		jsProt367 := factory366.GetProtocol(mbTrans364)
		argvalue0 := aurora.NewTaskQuery()
		err368 := argvalue0.Read(jsProt367)
		if err368 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.QueryRecovery(context.Background(), value0))
		fmt.Print("\n")
		break
	case "deleteRecoveryTasks":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DeleteRecoveryTasks requires 1 args")
			flag.Usage()
		}
		arg369 := flag.Arg(1)
		mbTrans370 := thrift.NewTMemoryBufferLen(len(arg369))
		defer mbTrans370.Close()
		_, err371 := mbTrans370.WriteString(arg369)
		if err371 != nil {
			Usage()
			return
		}
		factory372 := thrift.NewTJSONProtocolFactory()
		jsProt373 := factory372.GetProtocol(mbTrans370)
		argvalue0 := aurora.NewTaskQuery()
		err374 := argvalue0.Read(jsProt373)
		if err374 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DeleteRecoveryTasks(context.Background(), value0))
		fmt.Print("\n")
		break
	case "commitRecovery":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "CommitRecovery requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.CommitRecovery(context.Background()))
		fmt.Print("\n")
		break
	case "unloadRecovery":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "UnloadRecovery requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.UnloadRecovery(context.Background()))
		fmt.Print("\n")
		break
	case "startMaintenance":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartMaintenance requires 1 args")
			flag.Usage()
		}
		arg375 := flag.Arg(1)
		mbTrans376 := thrift.NewTMemoryBufferLen(len(arg375))
		defer mbTrans376.Close()
		_, err377 := mbTrans376.WriteString(arg375)
		if err377 != nil {
			Usage()
			return
		}
		factory378 := thrift.NewTJSONProtocolFactory()
		jsProt379 := factory378.GetProtocol(mbTrans376)
		argvalue0 := aurora.NewHosts()
		err380 := argvalue0.Read(jsProt379)
		if err380 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.StartMaintenance(context.Background(), value0))
		fmt.Print("\n")
		break
	case "drainHosts":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DrainHosts requires 1 args")
			flag.Usage()
		}
		arg381 := flag.Arg(1)
		mbTrans382 := thrift.NewTMemoryBufferLen(len(arg381))
		defer mbTrans382.Close()
		_, err383 := mbTrans382.WriteString(arg381)
		if err383 != nil {
			Usage()
			return
		}
		factory384 := thrift.NewTJSONProtocolFactory()
		jsProt385 := factory384.GetProtocol(mbTrans382)
		argvalue0 := aurora.NewHosts()
		err386 := argvalue0.Read(jsProt385)
		if err386 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DrainHosts(context.Background(), value0))
		fmt.Print("\n")
		break
	case "maintenanceStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "MaintenanceStatus requires 1 args")
			flag.Usage()
		}
		arg387 := flag.Arg(1)
		mbTrans388 := thrift.NewTMemoryBufferLen(len(arg387))
		defer mbTrans388.Close()
		_, err389 := mbTrans388.WriteString(arg387)
		if err389 != nil {
			Usage()
			return
		}
		factory390 := thrift.NewTJSONProtocolFactory()
		jsProt391 := factory390.GetProtocol(mbTrans388)
		argvalue0 := aurora.NewHosts()
		err392 := argvalue0.Read(jsProt391)
		if err392 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.MaintenanceStatus(context.Background(), value0))
		fmt.Print("\n")
		break
	case "endMaintenance":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "EndMaintenance requires 1 args")
			flag.Usage()
		}
		arg393 := flag.Arg(1)
		mbTrans394 := thrift.NewTMemoryBufferLen(len(arg393))
		defer mbTrans394.Close()
		_, err395 := mbTrans394.WriteString(arg393)
		if err395 != nil {
			Usage()
			return
		}
		factory396 := thrift.NewTJSONProtocolFactory()
		jsProt397 := factory396.GetProtocol(mbTrans394)
		argvalue0 := aurora.NewHosts()
		err398 := argvalue0.Read(jsProt397)
		if err398 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.EndMaintenance(context.Background(), value0))
		fmt.Print("\n")
		break
	case "slaDrainHosts":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "SlaDrainHosts requires 3 args")
			flag.Usage()
		}
		arg399 := flag.Arg(1)
		mbTrans400 := thrift.NewTMemoryBufferLen(len(arg399))
		defer mbTrans400.Close()
		_, err401 := mbTrans400.WriteString(arg399)
		if err401 != nil {
			Usage()
			return
		}
		factory402 := thrift.NewTJSONProtocolFactory()
		jsProt403 := factory402.GetProtocol(mbTrans400)
		argvalue0 := aurora.NewHosts()
		err404 := argvalue0.Read(jsProt403)
		if err404 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg405 := flag.Arg(2)
		mbTrans406 := thrift.NewTMemoryBufferLen(len(arg405))
		defer mbTrans406.Close()
		_, err407 := mbTrans406.WriteString(arg405)
		if err407 != nil {
			Usage()
			return
		}
		factory408 := thrift.NewTJSONProtocolFactory()
		jsProt409 := factory408.GetProtocol(mbTrans406)
		argvalue1 := aurora.NewSlaPolicy()
		err410 := argvalue1.Read(jsProt409)
		if err410 != nil {
			Usage()
			return
		}
		value1 := argvalue1
		argvalue2, err411 := (strconv.ParseInt(flag.Arg(3), 10, 64))
		if err411 != nil {
			Usage()
			return
		}
		value2 := argvalue2
		fmt.Print(client.SlaDrainHosts(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "snapshot":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Snapshot requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Snapshot(context.Background()))
		fmt.Print("\n")
		break
	case "triggerExplicitTaskReconciliation":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "TriggerExplicitTaskReconciliation requires 1 args")
			flag.Usage()
		}
		arg412 := flag.Arg(1)
		mbTrans413 := thrift.NewTMemoryBufferLen(len(arg412))
		defer mbTrans413.Close()
		_, err414 := mbTrans413.WriteString(arg412)
		if err414 != nil {
			Usage()
			return
		}
		factory415 := thrift.NewTJSONProtocolFactory()
		jsProt416 := factory415.GetProtocol(mbTrans413)
		argvalue0 := aurora.NewExplicitReconciliationSettings()
		err417 := argvalue0.Read(jsProt416)
		if err417 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.TriggerExplicitTaskReconciliation(context.Background(), value0))
		fmt.Print("\n")
		break
	case "triggerImplicitTaskReconciliation":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "TriggerImplicitTaskReconciliation requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.TriggerImplicitTaskReconciliation(context.Background()))
		fmt.Print("\n")
		break
	case "pruneTasks":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PruneTasks requires 1 args")
			flag.Usage()
		}
		arg418 := flag.Arg(1)
		mbTrans419 := thrift.NewTMemoryBufferLen(len(arg418))
		defer mbTrans419.Close()
		_, err420 := mbTrans419.WriteString(arg418)
		if err420 != nil {
			Usage()
			return
		}
		factory421 := thrift.NewTJSONProtocolFactory()
		jsProt422 := factory421.GetProtocol(mbTrans419)
		argvalue0 := aurora.NewTaskQuery()
		err423 := argvalue0.Read(jsProt422)
		if err423 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PruneTasks(context.Background(), value0))
		fmt.Print("\n")
		break
	case "createJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "CreateJob requires 1 args")
			flag.Usage()
		}
		arg424 := flag.Arg(1)
		mbTrans425 := thrift.NewTMemoryBufferLen(len(arg424))
		defer mbTrans425.Close()
		_, err426 := mbTrans425.WriteString(arg424)
		if err426 != nil {
			Usage()
			return
		}
		factory427 := thrift.NewTJSONProtocolFactory()
		jsProt428 := factory427.GetProtocol(mbTrans425)
		argvalue0 := aurora.NewJobConfiguration()
		err429 := argvalue0.Read(jsProt428)
		if err429 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.CreateJob(context.Background(), value0))
		fmt.Print("\n")
		break
	case "scheduleCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ScheduleCronJob requires 1 args")
			flag.Usage()
		}
		arg430 := flag.Arg(1)
		mbTrans431 := thrift.NewTMemoryBufferLen(len(arg430))
		defer mbTrans431.Close()
		_, err432 := mbTrans431.WriteString(arg430)
		if err432 != nil {
			Usage()
			return
		}
		factory433 := thrift.NewTJSONProtocolFactory()
		jsProt434 := factory433.GetProtocol(mbTrans431)
		argvalue0 := aurora.NewJobConfiguration()
		err435 := argvalue0.Read(jsProt434)
		if err435 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ScheduleCronJob(context.Background(), value0))
		fmt.Print("\n")
		break
	case "descheduleCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "DescheduleCronJob requires 1 args")
			flag.Usage()
		}
		arg436 := flag.Arg(1)
		mbTrans437 := thrift.NewTMemoryBufferLen(len(arg436))
		defer mbTrans437.Close()
		_, err438 := mbTrans437.WriteString(arg436)
		if err438 != nil {
			Usage()
			return
		}
		factory439 := thrift.NewTJSONProtocolFactory()
		jsProt440 := factory439.GetProtocol(mbTrans437)
		argvalue0 := aurora.NewJobKey()
		err441 := argvalue0.Read(jsProt440)
		if err441 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.DescheduleCronJob(context.Background(), value0))
		fmt.Print("\n")
		break
	case "startCronJob":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "StartCronJob requires 1 args")
			flag.Usage()
		}
		arg442 := flag.Arg(1)
		mbTrans443 := thrift.NewTMemoryBufferLen(len(arg442))
		defer mbTrans443.Close()
		_, err444 := mbTrans443.WriteString(arg442)
		if err444 != nil {
			Usage()
			return
		}
		factory445 := thrift.NewTJSONProtocolFactory()
		jsProt446 := factory445.GetProtocol(mbTrans443)
		argvalue0 := aurora.NewJobKey()
		err447 := argvalue0.Read(jsProt446)
		if err447 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.StartCronJob(context.Background(), value0))
		fmt.Print("\n")
		break
	case "restartShards":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RestartShards requires 2 args")
			flag.Usage()
		}
		arg448 := flag.Arg(1)
		mbTrans449 := thrift.NewTMemoryBufferLen(len(arg448))
		defer mbTrans449.Close()
		_, err450 := mbTrans449.WriteString(arg448)
		if err450 != nil {
			Usage()
			return
		}
		factory451 := thrift.NewTJSONProtocolFactory()
		jsProt452 := factory451.GetProtocol(mbTrans449)
		argvalue0 := aurora.NewJobKey()
		err453 := argvalue0.Read(jsProt452)
		if err453 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg454 := flag.Arg(2)
		mbTrans455 := thrift.NewTMemoryBufferLen(len(arg454))
		defer mbTrans455.Close()
		_, err456 := mbTrans455.WriteString(arg454)
		if err456 != nil {
			Usage()
			return
		}
		factory457 := thrift.NewTJSONProtocolFactory()
		jsProt458 := factory457.GetProtocol(mbTrans455)
		containerStruct1 := aurora.NewAuroraAdminRestartShardsArgs()
		err459 := containerStruct1.ReadField2(jsProt458)
		if err459 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.ShardIds
		value1 := argvalue1
		fmt.Print(client.RestartShards(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "killTasks":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "KillTasks requires 3 args")
			flag.Usage()
		}
		arg460 := flag.Arg(1)
		mbTrans461 := thrift.NewTMemoryBufferLen(len(arg460))
		defer mbTrans461.Close()
		_, err462 := mbTrans461.WriteString(arg460)
		if err462 != nil {
			Usage()
			return
		}
		factory463 := thrift.NewTJSONProtocolFactory()
		jsProt464 := factory463.GetProtocol(mbTrans461)
		argvalue0 := aurora.NewJobKey()
		err465 := argvalue0.Read(jsProt464)
		if err465 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		arg466 := flag.Arg(2)
		mbTrans467 := thrift.NewTMemoryBufferLen(len(arg466))
		defer mbTrans467.Close()
		_, err468 := mbTrans467.WriteString(arg466)
		if err468 != nil {
			Usage()
			return
		}
		factory469 := thrift.NewTJSONProtocolFactory()
		jsProt470 := factory469.GetProtocol(mbTrans467)
		containerStruct1 := aurora.NewAuroraAdminKillTasksArgs()
		err471 := containerStruct1.ReadField2(jsProt470)
		if err471 != nil {
			Usage()
			return
		}
		argvalue1 := containerStruct1.Instances
		value1 := argvalue1
		argvalue2 := flag.Arg(3)
		value2 := argvalue2
		fmt.Print(client.KillTasks(context.Background(), value0, value1, value2))
		fmt.Print("\n")
		break
	case "addInstances":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AddInstances requires 2 args")
			flag.Usage()
		}
		arg473 := flag.Arg(1)
		mbTrans474 := thrift.NewTMemoryBufferLen(len(arg473))
		defer mbTrans474.Close()
		_, err475 := mbTrans474.WriteString(arg473)
		if err475 != nil {
			Usage()
			return
		}
		factory476 := thrift.NewTJSONProtocolFactory()
		jsProt477 := factory476.GetProtocol(mbTrans474)
		argvalue0 := aurora.NewInstanceKey()
		err478 := argvalue0.Read(jsProt477)
		if err478 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err479 := (strconv.Atoi(flag.Arg(2)))
		if err479 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		fmt.Print(client.AddInstances(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "replaceCronTemplate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "ReplaceCronTemplate requires 1 args")
			flag.Usage()
		}
		arg480 := flag.Arg(1)
		mbTrans481 := thrift.NewTMemoryBufferLen(len(arg480))
		defer mbTrans481.Close()
		_, err482 := mbTrans481.WriteString(arg480)
		if err482 != nil {
			Usage()
			return
		}
		factory483 := thrift.NewTJSONProtocolFactory()
		jsProt484 := factory483.GetProtocol(mbTrans481)
		argvalue0 := aurora.NewJobConfiguration()
		err485 := argvalue0.Read(jsProt484)
		if err485 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.ReplaceCronTemplate(context.Background(), value0))
		fmt.Print("\n")
		break
	case "startJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "StartJobUpdate requires 2 args")
			flag.Usage()
		}
		arg486 := flag.Arg(1)
		mbTrans487 := thrift.NewTMemoryBufferLen(len(arg486))
		defer mbTrans487.Close()
		_, err488 := mbTrans487.WriteString(arg486)
		if err488 != nil {
			Usage()
			return
		}
		factory489 := thrift.NewTJSONProtocolFactory()
		jsProt490 := factory489.GetProtocol(mbTrans487)
		argvalue0 := aurora.NewJobUpdateRequest()
		err491 := argvalue0.Read(jsProt490)
		if err491 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.StartJobUpdate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "pauseJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "PauseJobUpdate requires 2 args")
			flag.Usage()
		}
		arg493 := flag.Arg(1)
		mbTrans494 := thrift.NewTMemoryBufferLen(len(arg493))
		defer mbTrans494.Close()
		_, err495 := mbTrans494.WriteString(arg493)
		if err495 != nil {
			Usage()
			return
		}
		factory496 := thrift.NewTJSONProtocolFactory()
		jsProt497 := factory496.GetProtocol(mbTrans494)
		argvalue0 := aurora.NewJobUpdateKey()
		err498 := argvalue0.Read(jsProt497)
		if err498 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.PauseJobUpdate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "resumeJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "ResumeJobUpdate requires 2 args")
			flag.Usage()
		}
		arg500 := flag.Arg(1)
		mbTrans501 := thrift.NewTMemoryBufferLen(len(arg500))
		defer mbTrans501.Close()
		_, err502 := mbTrans501.WriteString(arg500)
		if err502 != nil {
			Usage()
			return
		}
		factory503 := thrift.NewTJSONProtocolFactory()
		jsProt504 := factory503.GetProtocol(mbTrans501)
		argvalue0 := aurora.NewJobUpdateKey()
		err505 := argvalue0.Read(jsProt504)
		if err505 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.ResumeJobUpdate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "abortJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "AbortJobUpdate requires 2 args")
			flag.Usage()
		}
		arg507 := flag.Arg(1)
		mbTrans508 := thrift.NewTMemoryBufferLen(len(arg507))
		defer mbTrans508.Close()
		_, err509 := mbTrans508.WriteString(arg507)
		if err509 != nil {
			Usage()
			return
		}
		factory510 := thrift.NewTJSONProtocolFactory()
		jsProt511 := factory510.GetProtocol(mbTrans508)
		argvalue0 := aurora.NewJobUpdateKey()
		err512 := argvalue0.Read(jsProt511)
		if err512 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.AbortJobUpdate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "rollbackJobUpdate":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "RollbackJobUpdate requires 2 args")
			flag.Usage()
		}
		arg514 := flag.Arg(1)
		mbTrans515 := thrift.NewTMemoryBufferLen(len(arg514))
		defer mbTrans515.Close()
		_, err516 := mbTrans515.WriteString(arg514)
		if err516 != nil {
			Usage()
			return
		}
		factory517 := thrift.NewTJSONProtocolFactory()
		jsProt518 := factory517.GetProtocol(mbTrans515)
		argvalue0 := aurora.NewJobUpdateKey()
		err519 := argvalue0.Read(jsProt518)
		if err519 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.RollbackJobUpdate(context.Background(), value0, value1))
		fmt.Print("\n")
		break
	case "pulseJobUpdate":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PulseJobUpdate requires 1 args")
			flag.Usage()
		}
		arg521 := flag.Arg(1)
		mbTrans522 := thrift.NewTMemoryBufferLen(len(arg521))
		defer mbTrans522.Close()
		_, err523 := mbTrans522.WriteString(arg521)
		if err523 != nil {
			Usage()
			return
		}
		factory524 := thrift.NewTJSONProtocolFactory()
		jsProt525 := factory524.GetProtocol(mbTrans522)
		argvalue0 := aurora.NewJobUpdateKey()
		err526 := argvalue0.Read(jsProt525)
		if err526 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PulseJobUpdate(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getRoleSummary":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetRoleSummary requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetRoleSummary(context.Background()))
		fmt.Print("\n")
		break
	case "getJobSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobSummary requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobSummary(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTasksStatus":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksStatus requires 1 args")
			flag.Usage()
		}
		arg528 := flag.Arg(1)
		mbTrans529 := thrift.NewTMemoryBufferLen(len(arg528))
		defer mbTrans529.Close()
		_, err530 := mbTrans529.WriteString(arg528)
		if err530 != nil {
			Usage()
			return
		}
		factory531 := thrift.NewTJSONProtocolFactory()
		jsProt532 := factory531.GetProtocol(mbTrans529)
		argvalue0 := aurora.NewTaskQuery()
		err533 := argvalue0.Read(jsProt532)
		if err533 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksStatus(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTasksWithoutConfigs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetTasksWithoutConfigs requires 1 args")
			flag.Usage()
		}
		arg534 := flag.Arg(1)
		mbTrans535 := thrift.NewTMemoryBufferLen(len(arg534))
		defer mbTrans535.Close()
		_, err536 := mbTrans535.WriteString(arg534)
		if err536 != nil {
			Usage()
			return
		}
		factory537 := thrift.NewTJSONProtocolFactory()
		jsProt538 := factory537.GetProtocol(mbTrans535)
		argvalue0 := aurora.NewTaskQuery()
		err539 := argvalue0.Read(jsProt538)
		if err539 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetTasksWithoutConfigs(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getPendingReason":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetPendingReason requires 1 args")
			flag.Usage()
		}
		arg540 := flag.Arg(1)
		mbTrans541 := thrift.NewTMemoryBufferLen(len(arg540))
		defer mbTrans541.Close()
		_, err542 := mbTrans541.WriteString(arg540)
		if err542 != nil {
			Usage()
			return
		}
		factory543 := thrift.NewTJSONProtocolFactory()
		jsProt544 := factory543.GetProtocol(mbTrans541)
		argvalue0 := aurora.NewTaskQuery()
		err545 := argvalue0.Read(jsProt544)
		if err545 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetPendingReason(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getConfigSummary":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetConfigSummary requires 1 args")
			flag.Usage()
		}
		arg546 := flag.Arg(1)
		mbTrans547 := thrift.NewTMemoryBufferLen(len(arg546))
		defer mbTrans547.Close()
		_, err548 := mbTrans547.WriteString(arg546)
		if err548 != nil {
			Usage()
			return
		}
		factory549 := thrift.NewTJSONProtocolFactory()
		jsProt550 := factory549.GetProtocol(mbTrans547)
		argvalue0 := aurora.NewJobKey()
		err551 := argvalue0.Read(jsProt550)
		if err551 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetConfigSummary(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getJobs":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobs requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetJobs(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getQuota":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetQuota requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetQuota(context.Background(), value0))
		fmt.Print("\n")
		break
	case "populateJobConfig":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "PopulateJobConfig requires 1 args")
			flag.Usage()
		}
		arg554 := flag.Arg(1)
		mbTrans555 := thrift.NewTMemoryBufferLen(len(arg554))
		defer mbTrans555.Close()
		_, err556 := mbTrans555.WriteString(arg554)
		if err556 != nil {
			Usage()
			return
		}
		factory557 := thrift.NewTJSONProtocolFactory()
		jsProt558 := factory557.GetProtocol(mbTrans555)
		argvalue0 := aurora.NewJobConfiguration()
		err559 := argvalue0.Read(jsProt558)
		if err559 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.PopulateJobConfig(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getJobUpdateSummaries":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateSummaries requires 1 args")
			flag.Usage()
		}
		arg560 := flag.Arg(1)
		mbTrans561 := thrift.NewTMemoryBufferLen(len(arg560))
		defer mbTrans561.Close()
		_, err562 := mbTrans561.WriteString(arg560)
		if err562 != nil {
			Usage()
			return
		}
		factory563 := thrift.NewTJSONProtocolFactory()
		jsProt564 := factory563.GetProtocol(mbTrans561)
		argvalue0 := aurora.NewJobUpdateQuery()
		err565 := argvalue0.Read(jsProt564)
		if err565 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateSummaries(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDetails":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDetails requires 1 args")
			flag.Usage()
		}
		arg566 := flag.Arg(1)
		mbTrans567 := thrift.NewTMemoryBufferLen(len(arg566))
		defer mbTrans567.Close()
		_, err568 := mbTrans567.WriteString(arg566)
		if err568 != nil {
			Usage()
			return
		}
		factory569 := thrift.NewTJSONProtocolFactory()
		jsProt570 := factory569.GetProtocol(mbTrans567)
		argvalue0 := aurora.NewJobUpdateQuery()
		err571 := argvalue0.Read(jsProt570)
		if err571 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDetails(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getJobUpdateDiff":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetJobUpdateDiff requires 1 args")
			flag.Usage()
		}
		arg572 := flag.Arg(1)
		mbTrans573 := thrift.NewTMemoryBufferLen(len(arg572))
		defer mbTrans573.Close()
		_, err574 := mbTrans573.WriteString(arg572)
		if err574 != nil {
			Usage()
			return
		}
		factory575 := thrift.NewTJSONProtocolFactory()
		jsProt576 := factory575.GetProtocol(mbTrans573)
		argvalue0 := aurora.NewJobUpdateRequest()
		err577 := argvalue0.Read(jsProt576)
		if err577 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.GetJobUpdateDiff(context.Background(), value0))
		fmt.Print("\n")
		break
	case "getTierConfigs":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetTierConfigs requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetTierConfigs(context.Background()))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
