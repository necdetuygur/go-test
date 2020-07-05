(function(){
    if (location.protocol !== 'https:') {
        var httpsTestUrl = window.location.href.replace("http:", "https:") + "json";
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function() {
            if(this.status === 200){
                location.replace(`https:${location.href.substring(location.protocol.length)}`);
            }
        };
        xhr.open("GET", httpsTestUrl, true);
        xhr.send();
    }
})();
