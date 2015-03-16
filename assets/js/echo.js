var Label = React.createClass({displayName: "Label",
    render: function() {
        return (
            React.createElement("span", null, this.data.echoText)
        )
    }
});
React.render(
    React.createElement(Label, {echoText: "hoge"}),
    document.getElementById('label')
);