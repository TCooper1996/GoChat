package Models

const maxMessageSize = 512
const maxMessages = 40


type messageManager struct{
	messageArray [40]string

	// A pointer to the next index in the messageArray where the next message will be written
	mPointer int

}

func newMessageManager() *messageManager{
	return &messageManager{mPointer: 0}
}

func (m *messageManager) Push(message string){
	m.messageArray[m.mPointer] = message
}

