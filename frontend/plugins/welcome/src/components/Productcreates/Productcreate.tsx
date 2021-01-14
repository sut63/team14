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
  }),
);

/* interface CreateProduct {
  productname: string;
  numberofproduct: string;
  price: string;
  brand: number;
  typeproduct: number;
  personal: number;
} */

export default function CreateProductRecord() {
  const classes = useStyles();
  const api = new DefaultApi();
  
  const [products, setProduct] = React.useState<EntProduct[]>([]);
 
  const [brands, setBrands] = React.useState<EntBrand[]>([]);
  const [typeproducts, setTypeproducts] = React.useState<EntTypeproduct[]>([]);
  const [personals, setPersonals] = React.useState<EntPersonal[]>([]);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
  const [productname, setProductName] = useState(String);
  const [numberofproduct, setNumberofproduct] = useState(String);
  const [price, setPrice] = useState(String);

  const [brand, setBrand] = useState(Number);
  const [typeproduct, setTypeproduct] = useState(Number);
  const [personal, setPersonal] = useState(Number);

  useEffect(() => {
    const getBrands = async () => {
      const res = await api.listBrand({ limit: 10, offset: 0 });
      setLoading(false);
      setBrands(res);
      console.log(res);
    };
    getBrands();
  
    const getTypeproducts = async () => {
      const res = await api.listTypeproduct({ limit: 10, offset: 0 });
      setLoading(false);
      setTypeproducts(res);
      console.log(res);
    };
    getTypeproducts();
  
    const getPersonals = async () => {
      const res = await api.listPersonal({ limit: 10, offset: 0 });
      setLoading(false);
      setPersonals(res);
      console.log(res);
    };
    getPersonals();
  
  }, [loading]);
  
  
  const getProduct = async () => {
    const res = await api.listProduct({ limit: 10, offset: 0 });
    setProduct(res);
  };
  
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
  
  const PersonalhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setPersonal(event.target.value as number);
  };
  
  // create product
  const CreateProduct = async () => {
    const product = {
      productname : productname,
      numberofproduct : numberofproduct,
      price : price,
      brand : brand,
      typeproduct : typeproduct,
      personal : personal,
    };
    console.log(product);
    const res: any = await api.createProduct({ product : product });
    console.log("bruhhhhhhhhh");
    setStatus(true);
    
  };

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
            color="primary"
            > 
             ย้อนกลับ 
            </Button>
          </Link>
          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  <AlertTitle>บันทึกสำเร็จ</AlertTitle>
                </Alert>
              ) : (
                  <Alert severity="warning">
                    <AlertTitle>ไม่สามารถบันทึกข้อมูลได้ กรุณาลองใหม่อีกครั้ง</AlertTitle>
                  </Alert>
                )}
            </div>
          ) : null}
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
              style={{ width: 400 ,marginLeft:7,marginRight:-7}}
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
               style={{ width: 400 ,marginLeft:9,marginRight:-9}}
              InputProps={{
                startAdornment: (
                  <InputAdornment position="start">
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

            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <Typography gutterBottom  align="left">
                เจ้าหน้าที่ที่ทำการบันทึก :   
              <Typography variant="body1" gutterBottom> 
              <Select
                labelId="personal"
                id="personal"
                value={personal}
                onChange={PersonalhandleChange}
                style={{ width: 400 }}
              >
                {personals.map((item: EntPersonal) => (
                  <MenuItem value={item.id}>{item.personalname}</MenuItem>
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
                      CreateProduct();
                    }}
                      variant="contained"
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