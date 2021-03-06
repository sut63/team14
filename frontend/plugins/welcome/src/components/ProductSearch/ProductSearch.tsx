import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { EntProduct } from '../../api/models/EntProduct';

import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link, ContentHeader } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

import { Alert } from '@material-ui/lab';
import { styled } from '@material-ui/core/styles';
import { compose, spacing, palette, sizing, shadows   } from '@material-ui/system';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },


  }),
);
const Toast = Swal.mixin({
  //toast: true,
  position: 'center',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  showCloseButton: true,

});


export default function ComponentsTable() {

    //--------------------------
    ;

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  //---------------------------
  const [checkproductname, setProductnames] = useState(false);
  const [product, setProduct] = useState<EntProduct[]>([])

  //--------------------------
  const [piname, setpiname] = useState(String);
  const profile = { givenName: 'ระบบค้นหาข้อมูลอะไหล่คอมพิวเตอร์' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  //-------------------
  const productnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setpiname(event.target.value as string);

  };

  const cleardata = () => {
    setpiname("");
    setStatus(false)
    setProduct([]);
  }

  //---------------------
  const SearchProduct = async () => {
    setStatus(true);
    setAlert(true);
    const apiUrl = `http://localhost:8080/api/v1/searchproducts?product=${piname}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setStatus(true);
        setAlert(false);
        setProduct([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setStatus(true);
            setAlert(true);
            console.log(data.data)
            setProduct(data.data);
          }
        }
      });
    }


  return (

    <Page theme={pageTheme.tool}>
      <Header title={`Product System`} type="Computer Repair System" >
      <div>&nbsp;&nbsp;&nbsp;</div>
        <Button 
        style={{ marginLeft: 20 }} 
        href="/Producttables"
        variant="contained"  
        > 
        กลับหน้าตารางข้อมูลอะไหล่คอมพิวเตอร์ 
        </Button>

        <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        > 
        ออกจากระบบ 
        </Button>

      </Header>
      <Content>
      <ContentHeader title="ค้นหาข้อมูลบุคลากร">
          {status ? (
              <div>
                {alert ? (
                  <Alert severity="success">
                    ค้นหาสำเร็จ 
                  </Alert>
                )
                  : (
                    <Alert severity="warning" style={{ marginTop: 20 }}>
                      ไม่พบข้อมูลที่ค้นหา
                    </Alert>
                  )}
              </div>
            ) : null}
        </ContentHeader> 
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: 'linear-gradient(45deg, #CCCCCC 15%, #CCCCCC 120%)', height: 35 }}>
                  <h1 style={
                    {
                      //color: "#000000",
                      //borderRadius: 5,
                      //height: 18,
                      //padding: '0 30px',
                      fontSize: '20px',
                    }}>
                    ค้นหาข้อมูลอะไหล่คอมพิวเตอร์
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}><strong>กรุณากรอกชื่อสินค้า</strong></div>
                    <TextField
                      id="piname"
                      value={piname}
                      onChange={productnamehandlehange}
                      type="string"
                      size="small"

                      style={{ width: 250 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    SearchProduct();

                  }}
                  
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#5C9DC0", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    Search
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#DD0000", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 20px',

                      }
                    }>
                    Delete
            </h3>
                </Button>
              </Typography>
            </Paper>
          </Grid>
        </Grid>


        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              <TableContainer component={Paper}>
                <Table className={classes.table} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                    <TableCell align="center">No</TableCell>
                            <TableCell align="center">ชื่อสินค้า</TableCell>
                            <TableCell align="center">จำนวนสินค้า</TableCell>
                            <TableCell align="center">ราคา</TableCell>
                            <TableCell align="center">แบรนด์</TableCell>
                            <TableCell align="center">ประเภทของสินค้า</TableCell>
                            <TableCell align="center">เจ้าหน้าที่ที่ทำการบันทึก</TableCell>
                      
                    </TableRow>
                  </TableHead>
                <TableBody>
                {product.map((item: any) => (
                  <TableRow key={item.id}>
                    <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.Productname}</TableCell>
                                <TableCell align="center">{item.Amountofproduct}</TableCell>
                                <TableCell align="center">{item.Price}</TableCell>
                                <TableCell align="center">{item.edges?.Brand?.Brandname}</TableCell>
                                <TableCell align="center">{item.edges?.Typeproduct?.Typeproductname}</TableCell>
                                <TableCell align="center">{item.edges?.Personal?.Personalname}</TableCell>
                    
                  </TableRow>
                ))}
                </TableBody>
              </Table>
            </TableContainer>
            </Paper>
        </Grid>
        </Grid>
      </Content>
    </Page>
  );

}