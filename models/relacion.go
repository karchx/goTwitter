package models

/*Relacion modelo para relacionar un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionId string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
