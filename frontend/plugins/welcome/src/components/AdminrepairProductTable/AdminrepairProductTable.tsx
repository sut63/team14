import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';
import { EntProduct } from '../../api/models/EntProduct';

const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [products, setProducts] = useState<EntProduct[]>([]);
 
 useEffect(() => {
    const getProducts = async () => {
      const res:any = await api.listProduct({ limit: 10, offset: 0 });
      setLoading(true);
      setProducts(res);
      console.log(res);
    };
    getProducts();
  }, [loading]);
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">No.</TableCell>
           <TableCell align="center">ชื่อ อะไหล่</TableCell>
           <TableCell align="center">ประเภทอะไหล่</TableCell>
           <TableCell align="center">แบรนด์สินค้า</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {products.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.productname}</TableCell>
             <TableCell align="center">{item.edges.typeproduct.typeproductname}</TableCell>
             <TableCell align="center">{item.edges.brand.brandname}</TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
