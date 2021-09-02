# TO DO

- Remover todos os JSON do código
- Usar apenas os objetos "de verdade": objects.Player, utils.Vector, chat.Message
- Separar envio dos updates de mensagem e do player
  - Na primeira posição do byte array enviar um ID para identificar o tipo do objeto, se é mensagem, player ou outra coisa.