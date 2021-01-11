package ant

import (
	"fmt"
	"os"
)

func makeQEventStruct(gen *GenOut) (qEventServices []QEventService) {
	for i := 0; i < len(gen.Services); i++ {
		ser := gen.Services[i]
		var qEventsArr []QEvent

		for j := 0; j < len(ser.Methods); j++ {
			method := ser.Methods[j]

			qevent:= _buildQEvent(method,gen)

			qEventsArr = append(qEventsArr, qevent)
		}

		qeventSer := QEventService{
			ServiceName: ser.NameStriped,
			Events:      qEventsArr,
		}

		qEventServices = append(qEventServices, qeventSer)
	}

	PrettyPrint(qEventServices[2])
	return
}

func _buildQEvent(method MethodView, gen *GenOut) QEvent {
	qevent:= QEvent{
		EventName: method.MethodNameStriped,
		TagNum:    method.Pos + 5,
		Fields:    _buildQEventField(method.GoInTypeName, gen),
	}

	return qevent
}

var _msgViewMap = make(map[string]*MessageView)
func _buildQEventField(pramsMsgName string, gen *GenOut) (out []QEventPBFields) {
	// Fill the map
	if len(_msgViewMap) == 0 {
		for m := 0; m < len(gen.Messages); m++ {
			msg := gen.Messages[m]
			_msgViewMap[msg.Name] = &msg
		}
	}

	msg,found := _msgViewMap[pramsMsgName]
	if !found {
		panic(fmt.Sprintf("did not found pb param message: %s\n", pramsMsgName))
	}

	for g := 0; g < len(msg.Fields); g++ {
		msgFiled := msg.Fields[g]

		msgEventField := QEventPBFields{
			Name: msgFiled.FieldName    ,
			PBType:   msgFiled.TypeName,
			Repeated: msgFiled.Repeated,
		}

		out = append(out, msgEventField)
	}
	return
}


////////////////////////////////// bk ///////////////////////////////////
func makeQvent_bk(gen *GenOut) {
	var qventServices []QEventService

	for i := 0; i < len(gen.Services); i++ {
		ser := gen.Services[i]
		qeventSer := QEventService{
			ServiceName: ser.NameStriped,
			Events:      nil,
		}

		for j := 0; j < len(ser.Methods); j++ {
			method := ser.Methods[j]
			qevent:= QEvent{
				EventName: method.MethodNameStriped,
				TagNum:    method.Pos + 5,
				Fields:    nil,
			}

			// Add QEvents fileds for building message
			fondMsgParam := false
			for m := 0; m < len(gen.Messages); m++ {
				msg := gen.Messages[m]
				if msg.Name == method.GoInTypeName {
					fondMsgParam = true
					// Extract each pb message fields
					for g := 0; g < len(msg.Fields); g++ {
						msgFiled := msg.Fields[g]

						msgEventField := QEventPBFields{
							Name: msgFiled.FieldName    ,
							PBType:   msgFiled.TypeName,
							Repeated: msgFiled.Repeated,
						}

						qevent.Fields = append(qevent.Fields, msgEventField)

					}
					break // found the msg exit
				}
			}
			if !fondMsgParam {
				fmt.Printf("did not found pb param message: %s\n", method.GoInTypeName)
				os.Exit(1)
			}


			qeventSer.Events = append(qeventSer.Events, qevent)
		}
		qventServices = append(qventServices, qeventSer)
	}

	PrettyPrint(qventServices[2])
}