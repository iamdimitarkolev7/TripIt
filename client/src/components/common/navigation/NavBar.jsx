import * as React from 'react'
import { Link } from 'react-router-dom';
import { makeStyles } from '@mui/styles';
import AppBar from '@mui/material/AppBar';
import Typography from '@mui/material/Typography';
import Toolbar from '@mui/material/Toolbar';

const useStyles = makeStyles((theme) => ({
    navlinks: {
        marginLeft: theme.spacing(10),
        display: "flex",
      },
     logo: {
        flexGrow: "1",
        cursor: "pointer",
      },
      link: {
        textDecoration: "none",
        color: "white",
        fontSize: "20px",
        marginLeft: theme.spacing(20),
        "&:hover": {
          color: "black",
        },
      },
}))

const NavBar = ({isLoggedIn}) => {
    const classes = useStyles();

    return (
        <AppBar position="static">
            <Toolbar maxWidth="xl">
                <Typography className={classes.logo}>
                    Logo
                </Typography>
                <div className={classes.navlinks}>
                    <Link to="/" className={classes.link}>HOME</Link>
                    <Link to="/register" className={classes.link}>REGISTER</Link>
                    <Link to="/login" className={classes.link}>LOGIN</Link>
                    <Link to="/profile" className={classes.link}>PROFILE</Link>
                </div>
            </Toolbar>
        </AppBar>
    )
};

export default NavBar;