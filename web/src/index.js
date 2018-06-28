const axios = require('axios');
const echarts = require('echarts');
require('echarts-wordcloud');

// based on prepared DOM, initialize echarts instance
let chart = echarts.init(document.getElementById('main'));

document.querySelector("#queryForm").addEventListener("submit", (e) => {
    e.preventDefault();
    console.log("submit!");
    let query = document.querySelector('#queryForm input[name="query"]').value;
    let type_word = document.querySelector('#queryForm [name="type_word"]').value;

    console.log(type_word)

    let data = {query, type_word};

    axios.post('/api', data).then(response => {
        let keywords = response.data;
        console.log(keywords);

        let data = [];
        for (let name in keywords) {
            data.push({
                name: name,
                value: Math.sqrt(keywords[name])
            })
        }

        let option = {
            series: [{
                type: 'wordCloud',
                sizeRange: [10, 100],
                rotationRange: [-90, 90],
                rotationStep: 45,
                gridSize: 4,
                width: '100%',
                height: '100%',
                shape: 'circle',
                drawOutOfBound: false,
                textStyle: {
                    normal: {
                        fontWeight: 'bold',
                        color: function () {
                            return 'rgb(' + [
                                Math.round(Math.random() * 160),
                                Math.round(Math.random() * 160),
                                Math.round(Math.random() * 160)
                            ].join(',') + ')';
                        }
                    },
                    emphasis: {
                        color: 'red'
                    }
                },
                data: data.sort(function (a, b) {
                    return b.value - a.value;
                })
            }]
        };

        chart.setOption(option);
    });
});

chart.on('click', function (params) {
    console.log(params);
});

window.onresize = () => {
    chart.resize();
};
