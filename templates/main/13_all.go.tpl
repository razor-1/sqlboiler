{{- $alias := .Aliases.Table .Table.Name}}
{{- $schemaTable := .Table.Name | .SchemaTable}}
{{- $canSoftDelete := .Table.CanSoftDelete $.AutoColumns.Deleted }}
// {{$alias.UpPlural}} retrieves all the records using an executor.
func {{$alias.UpPlural}}(mods ...qm.QueryMod) {{$alias.DownSingular}}Query {
    {{if and .AddSoftDeletes $canSoftDelete -}}
    mods = append(mods, qm.From("{{$schemaTable}}"), qmhelper.WhereIsNull("{{$schemaTable}}.{{or $.AutoColumns.Deleted "deleted_at" | $.Quotes}}"))
    {{else -}}
    mods = append(mods, qm.From("{{$schemaTable}}"))
    {{end -}}

    q := NewQuery(mods...)
    if len(queries.GetSelect(q)) == 0 {
        if SpatialAsGeoJSON && len(SpatialTableColumns[TableNames.{{titleCase .Table.Name}}]) > 0 {
            // if we have a spatial column, then we need to enumerate each column to account for it.
            selectCols := make([]string, len({{$alias.DownSingular}}AllColumns))
            copy(selectCols, {{$alias.DownSingular}}AllColumns)
            for i, col := range selectCols {
                if isSpatialGeoJSON(TableNames.{{titleCase .Table.Name}}, col) {
                    selectCols[i] = fmt.Sprintf("ST_AsGeoJSON(`%s`) AS `%s`", col, col)
                }
            }
            queries.SetSelect(q, selectCols)
        } else {
            // otherwise, we can use the ".*" shortcut.
            queries.SetSelect(q, []string{"{{$schemaTable}}.*"})
        }
    }

    return {{$alias.DownSingular}}Query{q}
}
