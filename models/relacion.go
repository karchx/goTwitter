package models

/*Relacion modelo para relacionar un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usarioid" json:"usuarioId"`
	UsuarioRelacionId string `bson:"usariorelacionid" json:"usuarioRelacionId"`
}
