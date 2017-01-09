var Main = {
    _init: function() {
        Event._init();
        Guest._init();

        //init datetimepicker
        $(".datetimepicker-input").datetimepicker({
            format: "dd M yy hh:mm Z",
            locale: "ru"
        });
    }
};

var Event = {
    _init: function() {
        $('form').submit(function (e) {
            e.preventDefault();
        });
    },
    /**
     * @return {boolean}
     */
    HandleDelete : function($el) {
        var b = {"event_title": $el.data("title")};
        $.ajax({
            data : JSON.stringify(b),
            contentType : 'application/json',
            type: "POST",
            url: "/event/delete",
            success: function(m,s,r) {
                Alert.Do(m,s,r);
                $el.closest("tr").remove();
            },
            error: Alert.Do
        });
        return false;
    },
    /**
     * @return {boolean}
     */
    Edit : function(id) {
        var d = {"event_title": id};
        $.ajax({
            data : JSON.stringify(d),
            contentType : 'application/json',
            type: "POST",
            url: "/event/edit"
        });
        return false;
    },
    HandleEdit: function(form) {
        console.log("HandleEdit");
        return false;
    },
    EditTable: function(el) {
        console.log("EditTable");
        $("#dialog_table").modal();
        return false;
    },
    HandleAddTable: function(form) {
        console.log("HandleAddTable");
        return false;
    },
    HandleDeleteTable: function(idx) {
        console.log("HandleDeleteTable");
        return false;
    },
    handleCreateInited : false,
    HandleCreateInit : function() {
        if (Event.handleCreateInited) {
            return false;
        }
        $("#dialog").find(".btn-reset").bind({
            click: function (e) {
                $('form').trigger('reset');
            }
        });
        $("#dialog").find(".btn-list").bind({
            click: function(e) {
                window.location.replace("/event/list");
            }
        });
        Event.handleCreateInited = true;
    },
    /**
     * @return {boolean}
     */
    HandleCreate: function(form) {
        Event.HandleCreateInit();
        //Collect JSON for request
        var d = Utils.FormToJSON(form);
        //Do request
        $.ajax({
            data        : JSON.stringify(d),
            contentType : 'application/json',
            type        : "POST",
            url         : "/event/create",
            success     : function (s,m,r) {
                Dialog.Do(s,m,r, function($dg, data) {
                    $dg.find(".btn-edit").attr("href", "/event/edit/" + data)
                });
            },
            error       : Alert.Do
        });
        return false;
    }
};

var Guest = {
    _init: function() {
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
            data        : JSON.stringify(d),
            contentType : 'application/json',
            type        : "POST",
            url         : "/guest/create"
        });
        return false;
    },
    /**
     * @return {boolean}
     */
    HandleDelete : function (guestCode, eventTitle) {
        var d = {
            "guest_code"  : guestCode,
            "event_title" : eventTitle
        };
        $.ajax({
            data        : JSON.stringify(d),
            contentType : 'application/json',
            type        : "POST",
            url         : "/guest/delete"
        });
        return false;
    }
};

var Alert = {
    responseMsg: "Message",
    responseErr: "Error",

    self : this,
    BuildEmptyResponseJSON : function() {
        var d = {};
        d[this.responseMsg] = "";
        d[this.responseErr] = "";
        return d
    },
    Do: function (a, s, r) {
        if (typeof a === "object") {
            r = a
        }
        if (r.responseText == "") {
            var d = Alert.BuildEmptyResponseJSON();
            if (r.status == 200) {
                d[Alert.responseMsg] = "Good"
            } else {
                d[Alert.responseErr] = "Bad"
            }
        } else {
            d = JSON.parse(r.responseText);
        }
        console.log(d);
        var c = $(".container");
        var el  = d[Alert.responseErr] != ""
            ? Alert.ErrorEl(d[Alert.responseErr]) : Alert.SuccessEl(d[Alert.responseMsg]);
        c.find(".alert").remove();
        c.prepend(el);
    },
    SuccessEl: function (msg) {
        return $('<div class="alert alert-success fade in alert-dismissable" style="margin-top:18px;">' +
            '<a href="#" class="close" data-dismiss="alert" aria-label="close" title="close">×</a>' + msg + '</div>');
    },
    ErrorEl: function (msg) {
        return $('<div class="alert alert-danger fade in alert-dismissable" style="margin-top:18px;">' +
            '<a href="#" class="close" data-dismiss="alert" aria-label="close" title="close">×</a>' + msg + '</div>');
    }
};

var Dialog = {
    Do: function (a, s, r, fn) {
        var d = JSON.parse(r.responseText);
        var dg = $("#dialog");
        if (typeof fn == "function") {
            fn(dg, d["Message"]);
        }
        dg.modal();
    }
};

var Utils = {
    FormToJSON: function(form) {
        //Collect form data in to json format
        var f = form.serializeArray();
        var d = {};
        for (var i = 0; i < f.length; i++) {
            d[f[i].name] = f[i].value;
        }
        return d
    },
    InvokeErrorsFromArray: function(errors) {
        //If error occurred, show them all and stop process
        var c = $(".container");
        for (var j = 0;  j < errors.length; j++) {
            console.log(errors[j]);
            c.prepend(Alert.Error(errors[j]));
        }
    }
};