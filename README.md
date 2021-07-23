
# Panic and recover in Go make development less efficient

## 1. is use panic case

![img1](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723161926.png?raw=true)


### This is very risky behavior in the development stage.

![img2](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723161948.png?raw=true)
![img3](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723161955.png?raw=true)

### It is difficult to determine where the error is and it is frustrating at the front side.

### To solve this panic, go supports recover.
<br/><br/>

## 2. is use recover case
![img4](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723162549.png?raw=true)
![img5](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723162554.png?raw=true)
### Each recover function is required to expose information about errors.
![img6](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723165014.png?raw=true)
### It is possible to respond to errors, but it requires coercion and complexity of the code.
<br/><br/>

## 3. is use custom_errrorhandler case
![img8](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723163308.png?raw=true)
![img9](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723162613.png?raw=true)
### Basically, it reacts to errors in a handler like recover.
![img10](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723162618.png?raw=true)
### However, starting from the json code, a single error handler can return various errors.
![img11](https://github.com/sjy-dv/Go-CustomErrorHandler/blob/master/images/20210723163259.png?raw=true)
### You can handle the return of all errors flexibly with a single handler without having to write each recover individually.
