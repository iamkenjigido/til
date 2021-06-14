package structure

type Hoge struct {
	Name string
}

func (h *Hoge) HogeName(name string) string {
	h.Name = name
	return h.Name
}

func (h Hoge) NewHogeName(name string) string {
	h.Name = name
	return h.Name
}

type Fuga struct {
	Name string
}

func (f *Fuga) FugaName(name string) string {
	f.Name = name
	return f.Name
}

func (f Fuga) NewFugaName(name string) string {
	f.Name = name
	return f.Name
}

type Moge struct {
	Name string
}

func (m *Moge) MogeName(name string) string {
	m.Name = name
	return m.Name
}

func (m Moge) NewMogeName(name string) string {
	m.Name = name
	return m.Name
}