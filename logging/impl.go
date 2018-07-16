defer func() {
    {{.Recv.Name}}.logger.Log("service", "{{.Name}}", {{range $i, $e := .Params}}{{if gt $i 0}}"{{$e.Name}}", fmt.Sprintf("%+v",{{$e.Name}}), {{end}}{{end}})
}();
return {{.Recv.Name}}.{{.Name}}({{range .Params}}{{.Name}}, {{end}});
