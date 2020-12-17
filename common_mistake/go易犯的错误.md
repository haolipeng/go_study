一、for-range迭代slice或array
二、slice或array作为函数参数，函数内部修改其元素是否起作用？
三、将interface{}作为函数返回值类型的注意事项，只有interface的类型和值都为"nil"
interface判断时才能和nil比较，否则应附带提供error类型返回值
四、无效的switch判断
五、json、xml、protobuf序列化时，待序列化的结构体成员必须是大写
六、并非所有文本都是utf-8类型的
七、当把字符串转换为一个byte slice时(或反之),你就得到了一个原始数据的完整拷贝
八、slice易错点
    1） slice 破坏数据
    2） slice 隐藏数据（容量变大）
    3） 陈旧的slice
九、在go中，即使使用new()或make()函数来分配，变量的位置还是由编译器决定。
十、使用range来遍历数组或切片时，第一个返回的数值是下标索引，而不是数据或切片中的值