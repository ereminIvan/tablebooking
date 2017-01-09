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
        $("form").submit(function (e) {
            e.preventDefault();
        });
    },
    /**
     * @return {boolean}
     */
    HandleDelete : function($el) {
        var b = {"id": $el.data("id")};
        $.ajax({
            data : JSON.stringify(b),
            contentType : 'application/json',
            type: "POST",
            url: "/event/delete",
            success: function(m,s,r) {
                Alert.DoSuccess(m,s,r);
                $el.closest("tr").remove();
            },
            error: Alert.DoError
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
        console.info("HandleEdit");
        return false;
    },
    EditTable: function(el) {
        console.info("EditTable");
        $("#dialog_table").modal();
        return false;
    },
    HandleAddTable: function(form) {
        console.info("HandleAddTable");
        return false;
    },
    HandleDeleteTable: function(idx) {
        console.info("HandleDeleteTable");
        return false;
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
            success     : function (s, m, r) {
                Dialog.Do(s,m,r, function($dg, data) {
                    $dg.find(".btn-edit").attr("href", "/event/edit/" + data);
                });
            },
            error       : Alert.DoError
        });
        return false;
    },
    handleCreateInited : false,
    /**
     * @return {boolean}
     */
    HandleCreateInit : function() {
        if (Event.handleCreateInited) {
            return false;
        }
        var dg = $("#dialog");
        dg.find(".btn-reset").bind({
            click: function (e) {
                $('form').trigger('reset');
            }
        });
        dg.find(".btn-list").bind({
            click: function(e) {
                window.location.replace("/event/list");
            }
        });
        Event.handleCreateInited = true;
        return false;
    }
};

var Guest = {
    _init: function() {
        $('form').submit(function (evt) {
            evt.preventDefault();
        });
    },
    /**
     * Check given registration code
     * @return {boolean}
     */
    HandleCode : function(form) {
        $.post("/guest/code", {
            "guest_code" : form.val("guest_code")
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
    DoError: function (a, s, r) {
        var data = typeof a == "object" ? a : r;
        var d = JSON.parse(data.responseText);
        var c = $(".container");
        var el;
        c.find(".alert").remove();
        if (typeof  d["errors"] != "undefined") {
            for (var i = 0; i < d["errors"].length; i++) {
                el = Alert.ErrorEl(d["errors"][i]["title"]);
                c.prepend(el);
            }
        }
    },
    DoSuccess: function(a, s, r) {
        var data = typeof a == "object" ? a : r;
        var d = JSON.parse(data.responseText);
        var c = $(".container");
        var el;
        c.find(".alert").remove();
        if (typeof d["data"] != "undefined" && typeof d["data"]["message"] != "undefined") {
            el = Alert.SuccessEl(d["data"]["message"]);
            c.prepend(el);
        }
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
            fn(dg, d["data"]["id"]);
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