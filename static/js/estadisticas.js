const populateGraph = (info, ajaxR, animation) => {
    Highcharts.chart(info["id"], {
        chart: {
            type: 'bar'
        },
        title: {
            text: info["title"],
            align: 'left'
        },
        xAxis: {
            categories: ajaxR["categories"],
            title: {
                text: "Categorías"
            },
            gridLineWidth: 0,
            lineWidth: 2,
        },
        yAxis: {
            min: 0,
            title: {
                text: info["xlabel"],
                align: 'high'
            },
            labels: {
                overflow: 'justify'
            },
            gridLineWidth: 1,
            lineWidth: 2
        },
        plotOptions: {
            bar: {
                color: '#009999',
                animation: animation,
                borderRadius: '25%',
                dataLabels: {
                    enabled: true
                },
                groupPadding: 0,
            }
        },
        legend: {
            enabled: false
        },
        credits: {
            enabled: false
        },
        series: [{
            name:'',
            data: ajaxR["data"],
        }]
    });
};

let fetchAJAX = (info, animation) => {
    fetch(info["url"])
        .then((response) => {
            if (!response.ok)
                throw new Error("Network response was not ok");
            return response.json();
        })
        .then((ajaxResponse) => {
            populateGraph(info, ajaxResponse["data"], animation);
        })
        .catch((error) => {
            console.error(
                "There has been a problem with your fetch operation:",
                error
            );
        });
};

const dic_hinchas = {"id":"grafico_hinchas",   "titulo":"Estadísticas Hinchas",   "xlabel":"Hinchas por Deporte",     "url":url_grafico_hinchas};
const dic_artesan = {"id":"grafico_artesanos", "titulo":"Estadísticas Artesanos", "xlabel":"Artesanos por Artesanía", "url":url_grafico_artesanos};

setInterval(_ => {
    fetchAJAX(dic_hinchas, false);
    fetchAJAX(dic_artesan, false);
}, 2000);

fetchAJAX(dic_hinchas, true);
fetchAJAX(dic_artesan, true);


