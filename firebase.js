var PUBLICKEY = "<?PUBLICKEY?>";
var AUTHKEY = "<?AUTHKEY?>";
var TOPIC = "<?TOPIC?>";

var firebaseConfig = {
    apiKey: "<?APIKEY?>",
    authDomain: "fcm1-f0a24.firebaseapp.com",
    databaseURL: "https://fcm1-f0a24.firebaseio.com",
    projectId: "fcm1-f0a24",
    storageBucket: "fcm1-f0a24.appspot.com",
    messagingSenderId: "411358737390",
    appId: "1:411358737390:web:cbdb696be01b1ebb9d14ff"
};

firebase.initializeApp(firebaseConfig);

const messaging = firebase.messaging();

messaging.usePublicVapidKey(PUBLICKEY);

messaging.onTokenRefresh(() => {
    messaging.getToken().then((refreshedToken) => {
        subscribeTokenToTopic(refreshedToken, TOPIC);
    }).catch((err) => {
        console.log('Unable to retrieve refreshed token ', err);
    });
});

messaging.getToken().then((currentToken) => {
    if (currentToken) {
        subscribeTokenToTopic(currentToken, TOPIC);
    } else {
        console.log('No Instance ID token available. Request permission to generate one.');
    }
}).catch((err) => {
    console.log('An error occurred while retrieving token. ', err);
});

function requestPermission() {
    Notification.requestPermission().then((permission) => {
        if (permission === 'granted') {
            // console.log('Notification permission granted.');
        } else {
            console.log('Unable to Get permission to notify.');
        }
    });
}

function subscribeTokenToTopic(token, topic) {
    fetch('https://iid.googleapis.com/iid/v1/' + token + '/rel/topics/' + topic, {
        method: 'POST',
        headers: new Headers({
            'Authorization': "key=" + AUTHKEY,
        })
    }).then(response => {
        if (response.status < 200 || response.status >= 400) {
            throw 'subscribeTokenToTopic: ERR: ' + response.status + ' - ' + response.text();
        }
        // console.log('subscribeTokenToTopic: OK');
    }).catch(error => {
        console.error(error);
    })
}

window.addEventListener("load", requestPermission);