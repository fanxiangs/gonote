# chapter02_advfuncs
## [methods](./01-methods)

## [receivers](./02-receivers)

- 
  Pointers vs. Values

  不管你的method的receiver是指针类型还是非指针类型,都是可以通过指针/非指针类型进行调用的,编译器会帮你做类型转换。

- 在声明一个method的receiver该是指针还是非指针类型时,你需要考虑两方面

  第一方面是这个对象本身是不是特别大,如果声明为非指针变量时,调用会产生一次拷贝;

  第二方面是如果你用指针类型作为receiver,那么你一定要注意,这种指针类型指向的始终是一块内存地址,就算你对其进行了拷贝。


## [nonstructs](./03-nonstructs)

## [interfaces](./04-interfaces)
pirnter interface
[power](./04-interfaces/power): 带有pointer的interface

## [type assertion](./05-type-assertion)

## [empty interface](./06-empty-interface)

## [type switch](./07-type-switch)

## [promoted methods](./08-promoted-methods)

## [little refactor](./09-little-refactor)
接口分离
