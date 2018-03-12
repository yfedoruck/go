package data

type Data struct {
    Width int
    Height int
}

func (data *Data) IncreaseWidth (Width int) {
    data.Width += Width
}

func  (data *Data) IncreaseHeight (Height int) {
    data.Height += Height
}

func (data Data) Square () (int) {
    return data.Width * data.Height
}
