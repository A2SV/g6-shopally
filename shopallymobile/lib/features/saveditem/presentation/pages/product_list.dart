import 'package:flutter/material.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/detail_page.dart';
import 'package:shopallymobile/features/saveditem/presentation/pages/savedpage.dart';
import 'package:shopallymobile/features/saveditem/presentation/widgets/productlist/product_info.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import '../bloc/bloc/saved_product_bloc.dart';

class ProductListPage extends StatefulWidget {
  const ProductListPage({super.key});

  @override
  State<ProductListPage> createState() => _ProductListPageState();
}

class _ProductListPageState extends State<ProductListPage> {
  final List<dynamic> _savedItems = List.generate(5, (index) => 'Item $index');
  @override
  void initState() {
    super.initState();
    context.read<SavedProductBloc>().add(LoadSavedProductsEvent());
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Product List')),
      body: Center(
        child: Column(
          children: [
            ElevatedButton(onPressed: (){
              Navigator.push(context, MaterialPageRoute(builder: (context) => Savedpage()));
            },
            child: Text("SavedPage")),
            // ElevatedButton(onPressed: (){
            //   Navigator.push(context, MaterialPageRoute(builder: (context) => Demo()));
            // },
            // child: Text("Demo")),


        ElevatedButton(
          child: const Text('Show Modal'),
          onPressed: () {
            showModalBottomSheet(
              context: context,
              isScrollControlled: true,
              builder: (BuildContext context) {
                return Padding(
                  padding: const EdgeInsets.only(top: 15.0),
                  child: DraggableScrollableSheet(
                    expand: false,
                    builder: (context, scrollController) {
                      return Stack(
                        alignment: Alignment.topCenter,
                        children: [
                          Container(
                            margin: const EdgeInsets.only(top: 5),
                            decoration: const BoxDecoration(
                              color: Colors.white,
                              borderRadius: BorderRadius.only(
                                topLeft: Radius.circular(20),
                                topRight: Radius.circular(20),
                              ),
                            ),
                            child: CustomScrollView(
                              controller: scrollController,
                              slivers: <Widget>[
                                SliverAppBar(
                                  pinned: true,
                                  backgroundColor: Colors.white,
                                  automaticallyImplyLeading: false,
                                  elevation: 0,
                                  title: Row(
                                    mainAxisAlignment:
                                        MainAxisAlignment.spaceBetween,
                                    children: [
                                      IconButton(
                                        icon: const Icon(
                                          Icons.arrow_back_ios,
                                          color: Colors.black,
                                        ),
                                        onPressed: () {
                                          Navigator.of(context).pop();
                                        },
                                      ),
                                      const Text(
                                        'Product Search',
                                        style: TextStyle(
                                          fontSize: 18,
                                          fontWeight: FontWeight.bold,
                                          color: Colors.black,
                                        ),
                                      ),
                                      TextButton(
                                        onPressed: () {
                                          // Compare action
                                        },
                                        child: const Text('Compare'),
                                      ),
                                    ],
                                  ),
                                  bottom: PreferredSize(
                                    preferredSize: const Size.fromHeight(40),
                                    child: Align(
                                      alignment: Alignment.centerLeft,
                                      child: Padding(
                                        padding: const EdgeInsets.only(left: 16.0, bottom: 8.0),
                                        child: Text(
                                          'Book',
                                          style: Theme.of(context)
                                              .textTheme
                                              .titleLarge
                                              ?.copyWith(
                                                fontWeight: FontWeight.bold,
                                                color: Colors.black,
                                              ),
                                        ),
                                      ),
                                    ),
                                  ),
                                ),
                                SliverToBoxAdapter(
                                  child: Padding(
                                    padding: const EdgeInsets.all(16.0),
                                    child: Column(
                                      crossAxisAlignment:
                                          CrossAxisAlignment.start,
                                      children: const [
                                        SizedBox(height: 8),
                                        Text(
                                          'This is a description of the searched book. It can be a bit long and will wrap to multiple lines if needed.',
                                        ),
                                      ],
                                    ),
                                  ),
                                ),
                                 SliverPadding(
                  padding: const EdgeInsets.all(8.0),
                  sliver: SliverGrid(
                    gridDelegate:
                        const SliverGridDelegateWithFixedCrossAxisCount(
                          crossAxisCount: 2,
                          crossAxisSpacing: 2.0,
                          mainAxisSpacing: 8.0,
                          childAspectRatio: 0.55,
                        ),
                    delegate: SliverChildBuilderDelegate((context, index) {
                      final item = _savedItems[index];
                      return GestureDetector(
                        onTap: () {
                          Navigator.push(
                            context,
                            MaterialPageRoute(
                              builder: (context) => DetailPage(item: item),
                              
                            ),
                          );
                        },
                  
                        child: ProductInfo(id: index.toString()),
                      );
                    }, childCount: _savedItems.length),
                  ),
                ),
                              ],
                            ),
                          ),
                          Container(
                            width: 40,
                            height: 4,
                            margin: const EdgeInsets.only(top: 4),
                            decoration: BoxDecoration(
                              color: Colors.grey[300],
                              borderRadius: BorderRadius.circular(2),
                            ),
                          ),
                        ],
                      );
                    },
                  ),
                );
              },
            );
          },
        ),
      ]
      ),
      ),
    );
  }
}
