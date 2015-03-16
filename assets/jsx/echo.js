var Label = React.createClass({
    render: function() {
        return (
            <span>{this.data.echoText}</span>
        )
    }
});
React.render(
    <Label echoText="hoge" />,
    document.getElementById('label')
);