package Injector

type BeanFactory struct {
	beanMapper BeanMapper
}

func NewBeanFactory() *BeanFactory {
	return &BeanFactory{
		beanMapper: make(BeanMapper),
	}
}
