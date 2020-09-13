let store = (function()
{
    class Store
    {
        locStore = localStorage;
        listeners= {};
        constructor(options)
        {
            let l = this.locStore.getItem(this.name) || "{}";
            l = JSON.parse(l);
            for(let key in l)
                {
                    this[key] = l[key];
                }
        }
        set(obj, toLocStore, v)
        {
            let l = this.locStore.getItem(this.name) || "{}";
            l = JSON.parse(l);
            if(toLocStore)
            {
                for(let key in obj)
                {
                    l[key] = obj[key];
                    this[key] = obj[key];
                }
                this.locStore.setItem(this.name, JSON.stringify(l));
            } else 
            {
                for(let key in obj)
                {
                    this[key] = obj[key];
                    if(l[key]) this.locStore.setItem(this.name, JSON.stringify(l));
                }
            }
            for(let key in obj)
                {
                    if(this.listeners[key])
                    {
                        this.listeners[key].forEach(_ =>
                        {
                            _(obj[key], this);
                        });
                    }
                }
        };
        get(prop)
        {
            if(this[prop]) 
            {
                return utils.copyObj(this[prop]);
            } else 
            {
                let r = this.locStore.getItem(this.name) || "{}";
                r = JSON.parse(r);
                return r[prop];
            }
        };
        
        addListener(prop, callback)
        {
            callback ? "" : console.error("callback is undefined.");
            prop ? "" : console.error("property is undefined.");
            if(!this.listeners[prop]) this.listeners[prop] = [];
            if(this.listeners[prop].indexOf(prop) === -1) 
            {
                this.listeners[prop].push(callback);
            } else 
            {
                return "Этот обработчик уже назначен";
            }
        };
    }

    return new Store();
})()