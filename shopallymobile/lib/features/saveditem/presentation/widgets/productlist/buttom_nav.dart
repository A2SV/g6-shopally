// import 'package:flutter/material.dart';
// import 'package:shopallymobile/features/saveditem/presentation/pages/detail_page.dart';
// import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/product_info.dart';
// import 'package:shopallymobile/features/shopping_assistant/domain/entities/product_entity.dart';

// class Buttomnav extends StatefulWidget {
//   const Buttomnav({super.key, required this.products});

//   final List<ProductEntity> products;

//   @override
//   State<Buttomnav> createState() => _ButtomnavState();
// }

// class _ButtomnavState extends State<Buttomnav> {
//   // Dummy saved items list for demonstration
//   final List<ProductEntity> _savedItems = [];
//   @override
//   void initState() {
//     super.initState();
//     _savedItems.addAll(widget.products);
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Padding(
//       padding: const EdgeInsets.only(top: 0.0),
//       child: DraggableScrollableSheet(
//         initialChildSize: 1.0,
//         minChildSize: 1.0,
//         maxChildSize: 1.0,
//         expand: true,
//         builder: (context, scrollController) {
//           return Stack(
//             alignment: Alignment.topCenter,
//             children: [
//               Container(
//                 // Remove margin and borderRadius for full screen
//                 color: Colors.white,
//                 child: CustomScrollView(
//                   controller: scrollController,
//                   slivers: <Widget>[
//                     SliverAppBar(
//                       pinned: true,
//                       backgroundColor: Colors.white,
//                       automaticallyImplyLeading: false,
//                       elevation: 0,
//                       title: Row(
//                         mainAxisAlignment: MainAxisAlignment.spaceBetween,
//                         children: [
//                           IconButton(
//                             icon: const Icon(
//                               Icons.arrow_back_ios,
//                               color: Colors.black,
//                             ),
//                             onPressed: () {
//                               Navigator.of(context).pop();
//                             },
//                           ),
//                           const Text(
//                             'Product Search',
//                             style: TextStyle(
//                               fontSize: 18,
//                               fontWeight: FontWeight.bold,
//                               color: Colors.black,
//                             ),
//                           ),
//                           TextButton(
//                             onPressed: () {
//                               // Compare action
//                             },
//                             child: const Text('Compare'),
//                           ),
//                         ],
//                       ),
//                       bottom: PreferredSize(
//                         preferredSize: const Size.fromHeight(40),
//                         child: Align(
//                           alignment: Alignment.centerLeft,
//                           child: Padding(
//                             padding: const EdgeInsets.only(
//                               left: 16.0,
//                               bottom: 8.0,
//                             ),
//                             child: Text(
//                               'Book',
//                               style: Theme.of(context).textTheme.titleLarge
//                                   ?.copyWith(
//                                     fontWeight: FontWeight.bold,
//                                     color: Colors.black,
//                                   ),
//                             ),
//                           ),
//                         ),
//                       ),
//                     ),
//                     SliverToBoxAdapter(
//                       child: Padding(
//                         padding: const EdgeInsets.all(16.0),
//                         child: Column(
//                           crossAxisAlignment: CrossAxisAlignment.start,
//                           children: const [
//                             SizedBox(height: 8),
//                             Text(
//                               'This is a description of the searched book. It can be a bit long and will wrap to multiple lines if needed.',
//                             ),
//                           ],
//                         ),
//                       ),
//                     ),
//                     SliverPadding(
//                       padding: const EdgeInsets.all(8.0),
//                       sliver: SliverGrid(
//                         gridDelegate:
//                             const SliverGridDelegateWithFixedCrossAxisCount(
//                               crossAxisCount: 2,
//                               crossAxisSpacing: 2.0,
//                               mainAxisSpacing: 8.0,
//                               childAspectRatio: 0.55,
//                             ),
//                         delegate: SliverChildBuilderDelegate((context, index) {
//                           // final item = _savedItems[index]; // Removed unused variable
//                           return GestureDetector(
//                             onTap: () {
//                               Navigator.push(
//                                 context,
//                                 MaterialPageRoute(
//                                   builder: (context) => DetailPage(item: _savedItems[index]),
//                                 ),
//                               );
//                             },
//                             child: ProductInfo(
//                               id: _savedItems[index]['id'],
//                               title: _savedItems[index]['productName'],
//                               image: _savedItems[index]['imageUrl'],
//                               price: _savedItems[index]['priceUsd'],
//                               // minOrder: _savedItems[index]['sellerScore'].,
//                               rating: _savedItems[index]['productRating'] / 20,
//                             ),
//                           );
//                         }, childCount: _savedItems.length),
//                       ),
//                     ),
//                   ],
//                 ),
//               ),
//               Container(
//                 width: 40,
//                 height: 4,
//                 margin: const EdgeInsets.only(top: 4),
//                 decoration: BoxDecoration(
//                   color: Colors.grey[300],
//                   borderRadius: BorderRadius.circular(2),
//                 ),
//               ),
//             ],
//           );
//         },
//       ),
//     );
//   }
// }

