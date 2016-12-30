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
    HandlerCreate: function(form) {
        $.post("/event/create", {
            "event_title": form.val("event_title")
        }, function(data, status) {
            console.log(data, status);
        });
        return false;
    }
};

var Guest = {
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
        $.post("/guest/create", {
            "first_name": form.val("first_name"),
            "last_name":  form.val("last_name"),
            "is_vip":     form.val("is_vip")
        }, function(data, status) {
            console.log(data, status);
            if (status == 200) {
                $("body").html(data);
            } else {
                $("body").append(status)
            }
        });
        return false;
    }
};