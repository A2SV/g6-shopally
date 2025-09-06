import 'package:flutter/material.dart';
import 'package:shopallymobile/features/comparing/domain/Entity/comparison_entity.dart';

class Recommendation extends StatelessWidget {
  List<ComparisonEntity> comparisonEntities;
  Recommendation({super.key, required this.comparisonEntities});

  final List<Map<String, dynamic>> products = [
    {
      'name': 'Wireless Bluetooth Headphones',
      'pros': [
        'Excellent battery life',
        'Active noise cancellation',
        'Comfortable for long use',
        'Great value for money',
      ],
      'cons': [
        'Build quality could be better',
        'Limited color options',
        'No wireless charging',
      ],
      'score': 94.0,
    },
    {
      'name': 'Smart Fitness Tracker',
      'pros': [
        'Comprehensive health tracking',
        'Waterproof design',
        'Long battery life',
        'Detailed deep analysis',
      ],
      'cons': ['Higher price point', 'Small display', 'Limited app ecosystem'],
      'score': 56.23,
    },
    {
      'name': 'Portable Phone Charger',
      'pros': [
        'High capacity',
        'Fast charging support',
        'Multiple device charging',
        'LED battery indicator',
      ],
      'cons': ['Somewhat heavy', 'No wireless charging', 'Basic design'],
      'score': 70.3,
    },
  ];

  Widget _prosCons(List<Map<String, dynamic>> products) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Column(
        spacing: 20,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          for (var product in comparisonEntities)
            Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                SizedBox(
                  height: 40,
                  child: Text(
                    product.product.title,
                    maxLines: 1,
                    overflow: TextOverflow.ellipsis,
                    style: TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.w600,
                      color: Color.fromRGBO(38, 43, 50, 1),
                    ),
                  ),
                ),
                Row(
                  spacing: 5,
                  children: [
                    Icon(Icons.check_circle),
                    Text(
                      'Pros',
                      style: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.w500,
                        color: Color.fromRGBO(38, 43, 50, 1),
                      ),
                    ),
                  ],
                ),
                for (var pros in product.pros)
                  Row(
                    spacing: 5,
                    children: [
                      Icon(Icons.check, size: 12),
                      Expanded(
                        child: Text(
                          pros,
                          style: TextStyle(
                            fontSize: 14,
                            fontWeight: FontWeight.w400,
                            color: Color.fromRGBO(117, 123, 129, 1),
                          ),
                        ),
                      ),
                    ],
                  ),

                SizedBox(height: 10),

                Row(
                  spacing: 5,
                  children: [
                    Icon(Icons.cancel),
                    Text(
                      'Cons',
                      style: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.w500,
                        color: Color.fromRGBO(38, 43, 50, 1),
                      ),
                    ),
                  ],
                ),
                for (var cons in product.cons)
                  Row(
                    spacing: 5,
                    children: [
                      Icon(Icons.close, size: 12),
                      Expanded(
                        child: Text(
                          cons,
                          style: TextStyle(
                            fontSize: 14,
                            fontWeight: FontWeight.w400,
                            color: Color.fromRGBO(117, 123, 129, 1),
                          ),
                        ),
                      ),
                    ],
                  ),

                SizedBox(height: 10),
                Text(
                  'AI Score',
                  style: TextStyle(
                    fontSize: 12,
                    fontWeight: FontWeight.w500,
                    color: Color.fromRGBO(38, 43, 50, 1),
                  ),
                ),
                SizedBox(height: 5),
                LinearProgressIndicator(
                  value: product.product.aiMatchPercentage.round() / 100,
                  backgroundColor: Colors.grey.shade300,
                  valueColor: AlwaysStoppedAnimation<Color>(
                    product.isBest
                        ? Color.fromRGBO(255, 211, 0, 1)
                        : Color.fromRGBO(117, 123, 129, 1),
                  ),
                ),
                SizedBox(height: 5),
                Text(
                  '${product.product.aiMatchPercentage}% match',
                  style: TextStyle(
                    fontSize: 12,
                    fontWeight: FontWeight.w400,
                    color: Color.fromRGBO(117, 123, 129, 1),
                  ),
                ),
              ],
            ),
        ],
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(15.0),
      child: Container(
        decoration: const BoxDecoration(
          color: Colors.white,
          borderRadius: BorderRadius.all(Radius.circular(10)),
          boxShadow: [
            BoxShadow(
              color: Color.fromRGBO(0, 0, 0, 0.5),
              spreadRadius: 1,
              blurRadius: 10,
            ),
          ],
        ),
        child: Padding(
          padding: const EdgeInsets.all(15.0),
          child: Column(
            spacing: 20,
            children: [
              Column(children: [_prosCons(products)]),
            ],
          ),
        ),
      ),
    );
  }
}
