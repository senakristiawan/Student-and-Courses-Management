package course


import "fmt"
type CourseInterface interface {
    GetCourseByName(name string) (*Course, error)
    AddCourse(course Course) ([]Course, error)
}


type Course struct {
    ID   string
    Name string
}


var DummyCourses = []Course{
    {
        ID:   "1",
        Name: "Math",
    },
    {
        ID:   "2",
        Name: "English",
    },
}


func (c *Course) GetCourseByName(name string) (*Course, error) {
    // TODO: implementasi fungsi ini untuk mendapatkan course berdasarkan nama
    /*
        HINT :
        - Gunakan perulangan untuk mencari course yang sesuai dari slice DummyCourses
        - Jika tidak ditemukan course yang sesuai, kembalikan error dengan pesan "Course not found"
        - Jika ditemukan course yang sesuai, kembalikan course tersebut
    */
    // start_answer
    for i := range DummyCourses {
        if DummyCourses[i].Name == name {
            return &DummyCourses[i], nil
        }
    }
    // end_answer
    return nil, fmt.Errorf("Course not found") // TODO: replace this
}


func (c *Course) AddCourse(course Course) ([]Course, error) {
    // TODO: implementasi fungsi ini untuk menambahkan course baru
    /*
        HINT:
        - Gunakan append untuk menambahkan course baru ke slice DummyCourses
        - Jika course yang akan ditambahkan sudah ada di slice DummyCourses, kembalikan error dengan pesan "Course already exist"
        - Jika course yang akan ditambahkan tidak ada di slice DummyCourses, kembalikan slice DummyCourses yang sudah ditambahkan course baru
    */
    // start_answer
    for i := range DummyCourses {
        if DummyCourses[i].Name == course.Name {
            return nil, fmt.Errorf("Course already exist")
        }
    }
    DummyCourses = append(DummyCourses, course)
    return DummyCourses, nil // TODO: replace this
    //end_answer
}






mahasiswa.go
package mahasiswa


import (
    "final-project/course"
    "fmt"
)


type Mahasiswa struct {
    NIM     string
    Name    string
    Age     int
    Courses []course.Course
}


type MahasiswaInterface interface {
    AddMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error)
    DeleteMahasiswa(nim string) ([]Mahasiswa, error)
    GetMahasiswa(nim string) (*Mahasiswa, error)
    GetMahasiswaByCourse(courseName string) ([]Mahasiswa, error)
    GetMahasiswaList() ([]Mahasiswa, error)
    UpdateMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error)
}


var DummyMahasiswa = []Mahasiswa{
    {
        NIM:  "123",
        Name: "John Doe",
        Age:  20,
        Courses: []course.Course{
            {
                ID:   "1",
                Name: "Math",
            },
        },
    },
    {
        NIM:  "456",
        Name: "Jane Doe",
        Age:  20,
        Courses: []course.Course{
            {
                ID:   "1",
                Name: "Math",
            },
            {
                ID:   "2",
                Name: "English",
            },
        },
    },
}


func (m *Mahasiswa) GetMahasiswaList() ([]Mahasiswa, error) {
    if len(DummyMahasiswa) == 0 {
        return nil, fmt.Errorf("Tidak ada data mahasiswa")
    }
    return DummyMahasiswa, nil
}


func (m *Mahasiswa) UpdateMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error) {
    var mahasiswaIndex int = -1


    for i, m := range DummyMahasiswa {
        if m.NIM == mahasiswa.NIM {
            mahasiswaIndex = i
        }
    }


    if mahasiswaIndex == -1 {
        return nil, fmt.Errorf("Mahasiswa dengan NIM %s tidak ditemukan", mahasiswa.NIM)
    }


    DummyMahasiswa[mahasiswaIndex] = mahasiswa


    return DummyMahasiswa, nil
}


func (m *Mahasiswa) AddMahasiswa(mahasiswa Mahasiswa) ([]Mahasiswa, error) {
    //TODO: silahkan implementasikan fungsi ini untuk menambahkan mahasiswa
    /*
        Hint:
        - Gunakan slice DummyMahasiswa untuk menyimpan data mahasiswa
        - Gunakan slice DummyMahasiswa sebagai return value
        - Jika NIM sudah ada, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> sudah ada"
        - Jika NIM belum ada, tambahkan mahasiswa ke slice DummyMahasiswa dan kembalikan slice DummyMahasiswa
    */


    //start_answer
    for i := range DummyMahasiswa {
        if DummyMahasiswa[i].Name == mahasiswa.Name {
            return nil, fmt.Errorf("Mahasiswa dengan NIM %s sudah ada", mahasiswa.NIM)
        }
    }
    DummyMahasiswa = append(DummyMahasiswa, mahasiswa)
    return DummyMahasiswa, nil // TODO: replace this
    //end_answer
}


func (m *Mahasiswa) DeleteMahasiswa(nim string) ([]Mahasiswa, error) {
    //TODO: silahkan implementasikan fungsi ini untuk menghapus mahasiswa
    /*
        Hint:
        - Gunakan slice DummyMahasiswa sebagai sumber data
        - Gunakan slice DummyMahasiswa sebagai return value
        - Jika NIM tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> tidak ditemukan" dan nil sebagai return value
        - Jika NIM ditemukan, hapus mahasiswa dari slice DummyMahasiswa dan kembalikan slice DummyMahasiswa
        - Untuk menghapus mahasiswa dari slice, gunakan perulangan dan slice slicing (slice[start:end])
          - ex : slice = append(slice[:index], slice[index+1:]...)
    */


    //start_answer
    for i := range DummyMahasiswa {
        if DummyMahasiswa[i].NIM == nim {
            DummyMahasiswa = append(DummyMahasiswa[:i], DummyMahasiswa[i+1:]...)
            return DummyMahasiswa, nil
        }
    }
    // end_answer
    return nil, fmt.Errorf("Mahasiswa dengan NIM %s tidak ditemukan", nim) //TODO: replace this
}


func (m *Mahasiswa) GetMahasiswa(nim string) (*Mahasiswa, error) {
    //TODO: silahkan implementasikan fungsi ini untuk mendapatkan mahasiswa
    /*
        Hint:
        - Gunakan slice DummyMahasiswa sebagai sumber data
        - Jika NIM tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan NIM <NIM> tidak ditemukan" dan nil sebagai return value
        - Jika NIM ditemukan, kembalikan mahasiswa yang ditemukan
    */


    //start_answer
    for i := range DummyMahasiswa {
        if DummyMahasiswa[i].NIM == nim {
            return &DummyMahasiswa[i], nil
        }
    }
    //end_answer


    return nil, fmt.Errorf("Mahasiswa dengan NIM %s tidak ditemukan", nim) //TODO: replace this
}


func (m *Mahasiswa) GetMahasiswaByCourse(courseName string) ([]Mahasiswa, error) {
    //TODO: silahkan implementasikan fungsi ini untuk mendapatkan mahasiswa berdasarkan course
    /*
        Hint:
        - Gunakan slice DummyMahasiswa sebagai sumber data
        - Jika course tidak ditemukan, kembalikan error dengan pesan "Mahasiswa dengan course dengan nama <courseName> tidak ditemukan" dan nil sebagai return value dari slice mahasiswa
        - Jika course ditemukan, kembalikan mahasiswa yang mengambil course tersebut
    */
    var mahasiswaByCourse []Mahasiswa
    //start_answer


    for i := range DummyMahasiswa {
        for j := range DummyMahasiswa[i].Courses {
            if DummyMahasiswa[i].Courses[j].Name == courseName {
                mahasiswaByCourse = append(mahasiswaByCourse, DummyMahasiswa[i])
                return mahasiswaByCourse, nil
            }
        }
    }
    //end_answer
    return nil, fmt.Errorf("Mahasiswa dengan course dengan nama %s tidak ditemukan", courseName) //TODO: replace this
}
