require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");

$(() => {
    $(document).on('click', 'input[type="checkbox"]', function (event) {
        if($(this).prop("checked") == true){
            console.log("Checkbox is checked.", $(this).attr("id"));
        }
        else if($(this).prop("checked") == false){
            console.log("Checkbox is unchecked.");
        }
    });
});




