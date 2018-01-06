$(document).ready(function() {
    var phone = GetQueryString("phone");
    var cc = GetQueryString("cc");
    if (phone != null) {
        $('#text').val(phone);
        $("#cc").val(cc);
        GetNumberInfo(phone, cc);
    }
    $('#submit').click(function() {
        var phone = $('#text').val().trim();
        var cc = $("#cc").val().trim();
        if (phone != "") {
            GetNumberInfo(phone, cc);
        }
    });
    $("#text").on('keypress', function(e) {
        if (e.keyCode != 13) return;
        var phone = $('#text').val().trim();
        var cc = $("#cc").val().trim();
        if (phone != "") {
            GetNumberInfo(phone, cc);
        }
    });
});

function GetNumberInfo(phone, cc) {
    $('#editor_holder').html("<h4>loading...</h4>");
    $("#visual").html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/?phone="+encodeURIComponent(phone)+"&cc="+encodeURIComponent(cc), cache: false,
        success: function(result) {
            $('#editor_holder').jsonview(result);
            visual(result)
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}

function one(k, v) {
    return "<fieldset><legend class=\"label label-info left\">"+
        k+"</legend>"+v+"</fieldset>";
}

function visual(doc) {
    var html = "";
    for (k in doc) {
        html += one(k, doc[k]);
    }
    $("#visual").html(html);
}
