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
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';


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
  // toast: true,
  position: 'center',
  showConfirmButton: false,
  //timer: 3000,
  //timerProgressBar: true,
  showCloseButton: true,

});


export default function ComponentsTable() {

    //--------------------------
    ;

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checkproductname, setProductnames] = useState(false);
  const [product, setProduct] = useState<EntProduct[]>([])

  //--------------------------
  const [productname, setProductname] = useState(String);
  const profile = { givenName: 'ระบบค้นหาข้อมูลอะไหล่คอมพิวเตอร์' };
  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
    setSearch(false);
  }

  useEffect(() => {
    const getProducts = async () => {
      const res = await api.listProduct({ offset: 0 });
      setLoading(false);
      setProduct(res);
    };
    getProducts();
  }, [loading]);

  //-------------------
  const productnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setProductnames(false);
    setProductname(event.target.value as string);

  };

  const cleardata = () => {
    setProductname("");
    setSearch(false);
    setProductnames(false);
    setSearch(false);

  }
  //---------------------
  const checkresearch = async () => {
    var check = false;
    product.map(item => {
      if (productname != "") {
        if (item.productname?.includes(productname)) {
          setProductnames(true);
          alertMessage("success", "ค้นหาสำเร็จ");
          check = true;
        }
      }
    })
    if (!check) {
      alertMessage("error", "ไม่พบข้อมูลที่ค้นหา");
    }
    console.log(checkproductname)
    if (productname == "") {
      alertMessage("info", "แสดงข้อมูลอะไหล่คอมพิวเตอร์ทั้งหมดในระบบ");
    }
  };

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
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: 'linear-gradient(45deg, #CCCCCC 15%, #CCCCCC 120%)', height: 45 }}>
                  <h1 style={
                    {
                      //color: "#000000",
                      //borderRadius: 5,
                      //height: 18,
                      //padding: '0 30px',
                      fontSize: '25px',
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
                      id="productname"
                      value={productname}
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
                    checkresearch();
                    setSearch(true);

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
                        padding: '0 25px',

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
              {search ? (
                <div>
                  {  checkproductname ? (
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

                          {product.filter((filter: any) => filter.productname.includes(productname)).map((item: any) => (
                            <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.productname}</TableCell>
                                <TableCell align="center">{item.numberofproduct}</TableCell>
                                <TableCell align="center">{item.price}</TableCell>
                                <TableCell align="center">{item.edges?.brand?.brandname}</TableCell>
                                <TableCell align="center">{item.edges?.typeproduct?.typeproductname}</TableCell>
                                <TableCell align="center">{item.edges?.personal?.personalname}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                    : productname == "" ? (
                      <div>
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
                                    <TableCell align="center">{item.productname}</TableCell>
                                    <TableCell align="center">{item.numberofproduct}</TableCell>
                                    <TableCell align="center">{item.price}</TableCell>
                                    <TableCell align="center">{item.edges?.brand?.brandname}</TableCell>
                                    <TableCell align="center">{item.edges?.typeproduct?.typeproductname}</TableCell>
                                    <TableCell align="center">{item.edges?.personal?.personalname}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );

}