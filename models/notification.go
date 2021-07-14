package models

//Notification é a model de notificações
type Notification struct {
	ID               int64  `json:"id"`
	Titulo           string `json:"titulo_notificacao"`
	Mensagem         string `json:"mensagem"`
	Link             string `json:"link_notificacao"`
	Usuario          string `json:"usuario"`
	DataEnvio        string `json:"data_envio"`
	DataVisualizacao string `json:"data_visualizacao"`
	DataCadastro     string `json:"data_cadastro"`
	UsuarioCadastro  string `json:"usuario_cadastro"`
	DataAlteracao    string `json:"data_alteracao"`
	UsuarioAlteracao string `json:"usuario_alteracao"`
}