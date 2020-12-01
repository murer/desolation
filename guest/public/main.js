(function ($) {

    function generateResp(text) {
        // {"Name":"echo","Payload":"aaa"}
        var target = $('.msg-output-panel')
        target.html('')
        // var qrcode = new QRCode(target[0], {
        //     text: JSON.stringify(obj),
        //     width: 400,
        //     height: 400,
        //     colorDark : "#000000",
        //     colorLight : "#ffffff",
        //     correctLevel : QRCode.CorrectLevel.H
        // });

        var qr = new JSQR();

        var code = new qr.Code();
        code.encodeMode = code.ENCODE_MODE.BYTE;
        code.version = code.DEFAULT;
        code.errorCorrection = code.ERROR_CORRECTION.H;

        var input = new qr.Input();
        input.dataType = input.DATA_TYPE.TEXT;
        input.data = {
            "text": text
        };

        var matrix = new qr.Matrix(input, code);

        var canvas = document.createElement('canvas');
        canvas.setAttribute('width', matrix.pixelWidth);
        canvas.setAttribute('height', matrix.pixelWidth);
        canvas.getContext('2d').fillStyle = 'rgb(0,0,0)';
        matrix.draw(canvas, 0, 0);
        target[0].appendChild(canvas);

        console.log('generated')
    }

    $('form').submit(function (evt) {
        evt.preventDefault()
        var msgInput = $(this).find('[name=msg]').val()
        if (msgInput.startsWith('draw:')) {
            generateResp(msgInput.substring(5))
            return
        }
        $.ajax({
            type: 'POST',
            dataType: 'text',
            url: 'api/command',
            data: msgInput,
            contentType: 'text/plain',
            success: function (ret) {
                $('.msg-output-text').val(ret)
                generateResp(ret)
            }
        })
    })

    $(window).ready(function() {
        $('.msg-input').val('AgAAAAEAAAsxpGA')
        $('form').submit()
    })

    window.generateResp = generateResp

})(jQuery)