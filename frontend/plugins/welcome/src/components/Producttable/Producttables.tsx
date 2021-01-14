import React, { useState, useEffect } from 'react';
import { ContentHeader, Content, Header, Page, pageTheme } from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntProduct} from '../../api/models/EntProduct';
import HomeRoundedIcon from '@material-ui/icons/HomeRounded';
import ExitToAppIcon from '@material-ui/icons/ExitToApp';
import SettingsTwoToneIcon from '@material-ui/icons/SettingsTwoTone';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [products, setProducts] = useState<EntProduct[]>([]);
  const [loading, setLoading] = useState(true);
 
  // get products
  useEffect(() => {
    const getProducts = async () => {
      const res = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(true);
      setProducts(res);
      console.log(res);
    };
    getProducts();
  }, [loading]);

const deleteProducts = async (id: number) => {
 const res = await api.deleteProduct({ id: id });
 setLoading(true);
};
 
 return (
  <Page theme={pageTheme.tool}>
    <Header title={`เพิ่มข้อมูลอะไหล่คอมพิวเตอร์`} type="Computer Repair System" >
      <div>&nbsp;&nbsp;&nbsp;</div>
      <Button 
        style={{ marginLeft: 20 }} 
        href="/"
        variant="contained"  
        startIcon={<ExitToAppIcon/>}
        > 
        ออกจากระบบ 
        </Button>
    </Header>
    <Content>

    <ContentHeader title="ตารางแสดงผลข้อมูล">
      <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="secondary" href="/Productcreate" startIcon={<SettingsTwoToneIcon/>}> เพิ่มข้อมูลอะไหล่คอมพิวเตอร์ </Button>
    <div>&nbsp;&nbsp;&nbsp;</div>
    <Button variant="contained" color="primary" href="/Group14" startIcon={<HomeRoundedIcon/>}> กลับหน้าหลัก </Button>
      </ContentHeader>

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
         {products.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.productname}</TableCell>
             <TableCell align="center">{item.numberofproduct}</TableCell>
             <TableCell align="center">{item.price}</TableCell>
             <TableCell align="center">{item.edges?.brand?.brandname}</TableCell>
             <TableCell align="center">{item.edges?.typeproduct?.typeproductname}</TableCell>
             <TableCell align="center">{item.edges?.personal?.personalname}</TableCell>
             <TableCell align="center">
             <Button
                 onClick={() => {
                    deleteProducts(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 DELETE
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
   </Page>
 );
}