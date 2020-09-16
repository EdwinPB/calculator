require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {
    $(document).on('click', 'input[type="radio"]', function (event) {
        if($(this).prop("checked") == true){
            console.log("Checkbox is checked.", $(this).attr("id"));
            var landURL = "/calculators/show/?theme="+$(this).attr("id")
            $.ajax({
                url: landURL,
                type: "GET",
                dataType: 'json',
                data: {},
            }).done(function(data) {
                window.location = landURL
            }).fail(function(error) {
                window.location = landURL
                console.log(error)
            })
        }
        else if($(this).prop("checked") == false){
            console.log("Checkbox is unchecked.");
        }
    });
});




