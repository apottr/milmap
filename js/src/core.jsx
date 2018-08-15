import React, { PropTypes } from 'react'
import { Input, Card } from 'semantic-ui-react'
import {
	BrowserRouter as Router,
	Route,
	Link,
	Switch
} from 'react-router-dom'

const ConRouter = () => (
	<Router>
		<Switch>
			<Route exact path="/" component={Container} />
		</Switch>
	</Router>
)

class Container extends React.Component {
	render(){
		return(
			<h1>Hello World</h1>
		)
	}
}



export default ConRouter;
