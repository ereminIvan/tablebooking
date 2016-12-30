var Event = {
    Delete : function(id) {
        console.log(id);
    },
    Edit : function(id) {
        console.log(id);
    },
    /**
     * @return {boolean}
     */
    HandleCreate: function(form) {
        //Collect form data in json
        var f = form.serializeArray();
        var d = {};
        for (var i = 0; i < f.length; i++) {
            d[f[i].name] = f[i].value;
        }
        //Do request
        $.ajax({
            data : JSON.stringify(d),
            contentType : 'application/json',
            type: "POST",
            url: "/event/create"
        });
        return false;
    }
};

var Guest = {
    _init: function() {
        console.log("Init Guest");
        $('form').submit(function (evt) {
            evt.preventDefault();
        });
    },
    Code : function(code) {
        console.log(code)
    },
    /**
     * Check given registration code
     * @return {boolean}
     */
    HandleCode : function(form) {
        $.post("/guest/code", {
            "guest_code": form.val("guest_code")
        }, function(data, status) {
            console.log(data,status);
        });
        return false;

    },
    /**
     * Create guest with given params
     * @return {boolean}
     */
    HandleCreate : function (form) {
        //Collect form data in json
        var f = form.serializeArray();
        var d = {};
        for (var i = 0; i < f.length; i++) {
            if (f[1].name == "is_vip") {
                d[f[i].name] = true;
                continue;
            }
            d[f[i].name] = f[i].value;
        }
        //Do request
        $.ajax({
            data : JSON.stringify(d),
            contentType : 'application/json',
            type: "POST",
            url: "/guest/create"
        });
        return false;
    },
    /**
     * @return {boolean}
     */
    HandleDelete : function (guestCode, eventTitle) {
        var d = {
            "guest_code" : guestCode,
            "event_title" : eventTitle
        };
        $.ajax({
            data : JSON.stringify(d),
            contentType : 'application/json',
            type: "POST",
            url: "/guest/delete"
        })
        return false;
    }
};