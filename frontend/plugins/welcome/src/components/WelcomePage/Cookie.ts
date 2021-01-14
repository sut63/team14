class Cookies {

    SetCookie(name:any,value:any,days:any) {
        var expires = "";
        if (days) {
            var date = new Date();
            date.setTime(date.getTime() + (days*24*60*60*1000));
            expires = "; expires=" + date.toUTCString();
        }
        document.cookie = name + "=" + (value || "")  + expires + "; path=/";
    }
    
    GetCookie(name:any="email") {
        var nameEQ = name + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
        }
        return null;
    }

    GetID(id:any="id") {
        var idEQ = id + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(idEQ) == 0) return c.substring(idEQ.length,c.length);
        }
        return null;
    }

    GetName(id:any="name") {
        var idEQ = id + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(idEQ) == 0) return c.substring(idEQ.length,c.length);
        }
        return null;
    }

    GetRole(role:any="name") {
        var roleEQ = role + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(roleEQ) == 0) return c.substring(roleEQ.length,c.length);
        }
        return null;
    }

    CheckLogin(arr:Array<any>, email:any, password:any){
        var boo = false
        arr.forEach((item) => {
            if(item.email === email && item.password === password){
                boo = true                
            }
        });
        return boo
    }

    SetID(arr:Array<any>, email:any, password:any){
        var boo = 0
        arr.forEach((item) => {
            if(item.email === email && item.password === password){
                boo = item.id                
            }
        });
        return boo
    }

    SetName(arr:Array<any>, email:any, password:any){
        var boo = 0
        arr.forEach((item) => {
            if(item.email === email && item.password === password){
                boo = item.name                
            }
        });
        return boo
    }

    ClearCookie(name:any="email",id:any="id",role:any="name"){
        console.log("name in ClearCookie => "+name);
        console.log("id in ClearCookie => "+id);
        console.log("role in ClearCookie => "+role);
        const date = new Date();
        date.setTime(date.getTime() + (-1 * 24 * 60 * 60 * 1000));
        document.cookie = name+"=; expires="+date.toUTCString()+"; path=/";
        document.cookie = id+"=; expires="+date.toUTCString()+"; path=/";
        document.cookie = role+"=; expires="+date.toUTCString()+"; path=/";
    }
}
export {Cookies}