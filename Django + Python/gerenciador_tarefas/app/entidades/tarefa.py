class Tarefa():
    def __init__(self, titulo, descricao, data_expiracao, prioridade, usuario):
        self.__titulo = titulo
        self.__descricao = descricao
        self.__data_expiracao = data_expiracao
        self.__prioridade = prioridade
        self.__usuario = usuario


    @property
    def titulo(self):
        return self.__titulo

    @titulo.setter
    def titulo(self, titulo):
        self.__titulo = titulo

    @property
    def descricao(self):
        return self.descricao

    @descricao.setter
    def titulo(self, descricao):
        self.__descricao = descricao

    @property
    def data_expiracao(self):
        return self.__data_expiracao

    @data_expiracao.setter
    def data_expiracao(self, data_expiracao):
        self.__data_expiracao = data_expiracao

    @property
    def prioridade(self):
        return self.prioridade

    @prioridade.setter
    def prioridade(self, prioridade):
        self.__prioridade = prioridade

    @property
    def usuario(self):
        return self.usuario

    @usuario.setter
    def usuario(self, usuario):
        self.__usuario = usuario