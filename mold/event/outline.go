package event

func (e *Event) Outline() (outline Event) {
	outline = *e
	outline.Raw = ""
	return
}
