import 'package:flutter/material.dart';

import 'package:shopallymobile/features/comparing/domain/Entity/comparison_result_entity.dart';
import 'package:shopallymobile/features/comparing/presentation/widget/recommendation.dart';

import '../widget/comparison_table.dart';
import '../widget/over_all_comparison_widget.dart';

class ComparisonResultPage extends StatelessWidget {
  final ComparisonResultEntity comparisonResultEntity;
  const ComparisonResultPage({super.key, required this.comparisonResultEntity});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: BackButton(
          onPressed: () => Navigator.pop(context),
        ),
        backgroundColor: Theme.of(context).scaffoldBackgroundColor,
        elevation: 0,
      ),
      body: SafeArea(
        child: SingleChildScrollView(
          child: Column(
            children: [
              ComparisonTableWidget(
                comparisonEntity: comparisonResultEntity.comparisonEntity,
              ),
              OverAllComparisonWidget(
                overAllComparisonEntity:
                    comparisonResultEntity.overAllComparisonEntity,
              ),
              Recommendation(
                comparisonEntities: comparisonResultEntity.comparisonEntity,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
