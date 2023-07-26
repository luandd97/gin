lc = $(subst A,a,$(subst B,b,$(subst C,c,$(subst D,d,$(subst E,e,$(subst F,f,$(subst G,g,$(subst H,h,$(subst I,i,$(subst J,j,$(subst K,k,$(subst L,l,$(subst M,m,$(subst N,n,$(subst O,o,$(subst P,p,$(subst Q,q,$(subst R,r,$(subst S,s,$(subst T,t,$(subst U,u,$(subst V,v,$(subst W,w,$(subst X,x,$(subst Y,y,$(subst Z,z,$1))))))))))))))))))))))))))
crm:
	touch controllers/$(arg)Controller.go 
	echo "package controllers\n\ntype $(arg)Controller interface{}\n\ntype $(call lc ,$(arg))Controller struct{}\n\nfunc New$(arg)Controller() $(arg)Controller {\n	return &$(call lc ,$(arg))Controller{}\n}" > controllers/$(arg)Controller.go 
	touch entities/$(arg).go 
	echo "package entities\n\ntype $(arg) struct{}" > entities/$(arg).go 
	touch repositories/$(arg)Repository.go
	echo "package repositories\n\ntype $(arg)Repository interface{}\n\ntype $(call lc ,$(arg))Repository struct{}\n\nfunc New$(arg)Repository() $(arg)Repository {\n	return &$(call lc ,$(arg))Repository{}\n}" > repositories/$(arg)Repository.go  


m:
	touch entities/$(arg).go 
	echo "package entities\n\ntype $(arg) struct{}" > entities/$(arg).go 

c:
	touch controllers/$(arg)Controller.go 
	echo "package controllers\n\ntype $(arg)Controller interface{}\n\ntype $(call lc ,$(arg))Controller struct{}\n\nfunc New$(arg)Controller() $(arg)Controller {\n	return &$(call lc ,$(arg))Controller{}\n}" > controllers/$(arg)Controller.go 

r:
	touch repositories/$(arg)Repository.go
	echo "package repositories\n\ntype $(arg)Repository interface{}\n\ntype $(call lc ,$(arg))Connection struct{}\n\nfunc New$(arg)Repository() $(arg)Repository {\n	return &$(call lc ,$(arg))Connection{}\n}" > repositories/$(arg)Repository.go  

s:
	touch services/$(arg)Service.go
	echo "package services\n\ntype $(arg)Service interface{}\n\ntype $(call lc ,$(arg))Service struct{}\n\nfunc New$(arg)Service() $(arg)Service {\n	return &$(call lc ,$(arg))Service{}\n}" > services/$(arg)Service.go  