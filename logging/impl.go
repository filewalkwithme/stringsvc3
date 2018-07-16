defer func() {
    {{.Recv.Name}}.logger.Log("service", "{{.Name}}", {{range $i, $e := .Params}}"{{$e.Name}}", fmt.Sprintf("%+v",{{$e.Name}}), {{end}})
}();
return {{.Recv.Name}}.{{.Name}}({{range .Params}}{{.Name}}, {{end}});
