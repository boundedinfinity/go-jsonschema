package jsonschema

// func (t *System) Load(paths ...string) error {
// 	for _, path := range paths {
// 		if err := fileutil.WalkFilePaths(path, t.hasExtention, t.walkProcess); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *System) hasExtention(path string) bool {
// 	for _, ext := range t.Extentions {
// 		if strings.HasSuffix(path, ext) {
// 			return true
// 		}
// 	}

// 	return false
// }

// func (t *System) walkProcess(path string) error {
// 	var mt mime_type.MimeType

// 	if err := t.detechMimeType(path, &mt); err != nil {
// 		return err
// 	}

// 	bs, err := os.ReadFile(path)

// 	if err != nil {
// 		return err
// 	}

// 	ss := []JsonSchema{}

// 	if err := t.Unmarshal(&ss, mt, bs); err != nil {
// 		return err
// 	}

// 	for i, s := range ss {
// 		key := fmt.Sprintf("%v:%v", path, i)

// 		if s.Id.Empty() {
// 			return ErrIdEmptyf(key)
// 		}

// 		t.SourceMap[key] = s.Id.Get()

// 		if s.system == nil {
// 			s.system = t
// 		}

// 		t.IdMap[s.Id.Get()] = s
// 	}

// 	return nil
// }

// func (t *System) detechMimeType(path string, mt *mime_type.MimeType) error {
// 	ext := filepath.Ext(path)
// 	temp, err := file_extention.Extention2MimeType(ext)

// 	if err != nil {
// 		return err
// 	}

// 	*mt = temp

// 	return nil
// }
