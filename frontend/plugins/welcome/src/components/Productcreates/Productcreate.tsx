import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle} from '@material-ui/lab';
import InputAdornment from '@material-ui/core/InputAdornment';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { Cookies } from '../WelcomePage/Cookie'
import Swal from 'sweetalert2';
import Product from '.';
import SaveIcon from '@material-ui/icons/Save';
import ArrowBackIosIcon from '@material-ui/icons/ArrowBackIos';
import MonetizationOnIcon from '@material-ui/icons/MonetizationOn';
import LocalOfferIcon from '@material-ui/icons/LocalOffer';

import { DefaultApi } from '../../api/apis';
import { EntBrand } from '../../api/models/EntBrand'; // import interface Brand
import { EntTypeproduct } from '../../api/models/EntTypeproduct'; // import interface Typeproduct
import { EntPersonal } from '../../api/models/EntPersonal'; // import interface Personal
import { EntProduct } from '../../api';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },
    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    select: {
      width: 500 ,
      marginLeft:7,
      marginRight:-7,
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    pa: {
      marginTop: theme.spacing(2),
    },
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

var ck = new Cookies()
var cookieName = ck.GetCookie()
var cookieID = ck.GetID()

export default function CreateProductRecord() {
  const classes = useStyles();
  const http = new DefaultApi();
  
  const [products, setProduct] = React.useState<EntProduct[]>([]);
  const [brands, setBrands] = React.useState<EntBrand[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);
  const [loading, setLoading] = useState(true);
 
  const [productname, setProductName] = useState(String);
  const [numberofproduct, setNumberofproduct] = useState(String);
  const [price, setPrice] = useState(String);

  const [brand, setBrand] = useState(Number);
  const [typeproduct, setTypeproduct] = useState(Number);
  const [personal, setPersonal] = useState(Number);

  useEffect(() => {
    const getBrands = async () => {
      const b = await http.listBrand({ limit: 10, offset: 0 });
      setLoading(false);
      setBrands(b);
    };
    getBrands();
  
    const getTypeproducts = async () => {
      const t = await http.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(t);
    };
    getTypeproducts();
  
    const getPersonals = async () => {
      const p = await http.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(p);
    };
    getPersonals();
  
  }, [loading]);
  
  
  /*const getProduct = async () => {
    const res = await http.listProduct({ limit: 10, offset: 0 });
    setProduct(res);
  };*/
  
  const handleproductnameChange = (event: any) => {
    setProductName(event.target.value as string);
  };
  
  const handlenumberofproductChange = (event: any) => {
    setNumberofproduct(event.target.value as string);
  };

  const handlepriceChange = (event: any) => {
    setPrice(event.target.value as string);
  };
  
  const BrandhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBrand(event.target.value as number);
  };
  
  const TypeproducthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setTypeproduct(event.target.value as number);
  };
  
  let p = Number(personal)

  const product = {
   productname : productname,
   numberofproduct : numberofproduct,
   price : price,
   brand : brand,
   personal : Number(cookieID),
   typeproduct : typeproduct,
   };
 
   const alertMessage = (icon: any, title: any) => {
     Toast.fire({
       icon: icon,
       title: title,
     });
   }
 
   const checkCaseSaveError = (field: string) => {
     switch(field) {
       case 'Productname':
         alertMessage("error","กรุณากรอกชื่อสินค้า");
         return;
       case 'Numberofproduct':
         alertMessage("error","กรุณากรอกจำนวนสินค้า");
         return;
       case 'Price':
         alertMessage("error","รูปแบบของราคาสินค้าไม่ถูกต้อง กรุณาตรวจสอบอีกครั้ง");
         return;
       default:
         alertMessage("error","บันทึกข้อมูลไม่สำเร็จ");
         return;
     }
   }
 
   console.log(product)
 function save() {
   const apiUrl = 'http://localhost:8080/api/v1/products';
   const requestOptions = {
     method: 'POST',
     headers: { 'Content-Type': 'application/json' },
     body: JSON.stringify(product),
   };
 
   console.log(product); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console
 
   fetch(apiUrl, requestOptions)
     .then(response => response.json())
     .then(data => {
       console.log(data);
       if (data.status == true) {
         //clear();
         Toast.fire({
           icon: 'success',
           title: 'บันทึกข้อมูลสำเร็จ',
 
         });//window.setTimeout(function(){location.reload()},8000);
       } else {
         checkCaseSaveError(data.error.Name)
       }
     });
 
     
     const timer = setTimeout(() => {
       setStatus(false);
     }, 1000);
   };

  /*// create product
  const CreateProduct = async () => {
    if ((productname != null) && (productname != "") && (numberofproduct != null) && (numberofproduct != "") && (price != null) && (price != "") && (brand != null) && (typeproduct != null)  && (personal != null) ) {
    
      const product = {
      productname : productname,
      numberofproduct : numberofproduct,
      price : price,
      brand : brand,
      typeproduct : typeproduct,
      personal : Number(cookieID),
    };
    console.log(products);
    const res: any = await http.createProduct({ product: product });
   
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } 
  }
    else {
      setStatus(true);
      setAlert(false);
    }
};*/

  return (
    <Page theme={pageTheme.tool}>

      <Header
        title={`เพิ่มอะไหล่คอมพิวเตอร์`}> 
      </Header>
      <Content>
        <ContentHeader title="กรอกข้อมูลอะไหล่คอมพิวเตอร์">
        <div>
            <Link component={RouterLink} to="Group14/">
            <Button 
            style={{ marginLeft: 20 }} 
            component={RouterLink} 
            to="/Producttables" 
            variant="contained"
            startIcon={<ArrowBackIosIcon />}
            color="primary"
            > 
             ย้อนกลับ 
            </Button>
          </Link>
          </div>
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <FormControl
              variant="outlined"
            >
               <div className={classes.paper}><strong>ชื่อสินค้า</strong></div>
              <TextField className={classes.textField}
               style={{ width: 400 ,marginLeft:9,marginRight:-9}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <LocalOfferIcon />
                  </InputAdornment>
                ),
              }}
                id="productname"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={productname}
                onChange={handleproductnameChange}
        />

            <div className={classes.paper}><strong>จำนวนสินค้า</strong></div>
              <TextField className={classes.textField}
              style={{ width: 400 ,marginLeft:9,marginRight:-9}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                  </InputAdornment>
                ),
              }}
                id="numberofproduct"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={numberofproduct}
                onChange={handlenumberofproductChange}
              />

            <div className={classes.paper}><strong>ราคา</strong></div>
              <TextField className={classes.textField}
               style={{ width: 400 ,marginLeft:7,marginRight:-7}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
                    <MonetizationOnIcon />
                  </InputAdornment>
                ),
              }}
                id="price"
                label=""
                variant="standard"
                color="secondary"
                type="string"
                size="medium"
                value={price}
                onChange={handlepriceChange}
            />
           
            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography gutterBottom  align="left">
                แบรนด์ :   
              <Typography variant="body1" gutterBottom> 
              <Select
                labelId="brand"
                id="brand"
                value={brand}
                onChange={BrandhandleChange}
                style={{ width: 400 }}
              >
                {brands.map((item: EntBrand) => (
                  <MenuItem value={item.id}>{item.brandname}</MenuItem>
                ))}
              </Select>
              </Typography>
                </Typography>
            </FormControl>
            </div>

            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography gutterBottom  align="left">
                ประเภทสินค้า :   
              <Typography variant="body1" gutterBottom> 
              <Select
                labelId="typeproduct"
                id="typeproduct"
                value={typeproduct}
                onChange={TypeproducthandleChange}
                style={{ width: 400 }}
              >
                {typeproducts.map((item: EntTypeproduct) => (
                  <MenuItem value={item.id}>{item.typeproductname}</MenuItem>
                ))}
              </Select>
              </Typography>
                </Typography>
            </FormControl>
            </div>

           

            <div className={classes.margin}>
                <Typography variant="h6" gutterBottom  align="center">
                  <Button
                    onClick={() => {
                      save();
                    }}
                      variant="contained"
                      startIcon={<SaveIcon />}
                      color="primary"
                  >
                    Submit
                  </Button>
                </Typography>
              </div>
              
              </FormControl>
          </form>
        </div>
      </Content>
    </Page>
  );
} 