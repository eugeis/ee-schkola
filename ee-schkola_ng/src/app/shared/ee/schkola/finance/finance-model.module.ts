import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {FinanceRoutingModules} from './finance-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {FinanceViewComponent} from './components/view/finance-module-view.component';
import {ExpenseViewComponent} from '@schkola/finance/expense/components/view/expense-entity-view.component';
import {ExpenseListComponent} from '@schkola/finance/expense/components/list/expense-entity-list.component';
import {ExpenseFormComponent} from '@schkola/finance/expense/components/form/expense-form.component';
import {ExpensePurposeViewComponent} from '@schkola/finance/expensepurpose/components/view/expensepurpose-entity-view.component';
import {ExpensePurposeListComponent} from '@schkola/finance/expensepurpose/components/list/expensepurpose-entity-list.component';
import {ExpensePurposeFormComponent} from '@schkola/finance/expensepurpose/components/form/expensepurpose-form.component';
import {FeeViewComponent} from '@schkola/finance/fee/components/view/fee-entity-view.component';
import {FeeListComponent} from '@schkola/finance/fee/components/list/fee-entity-list.component';
import {FeeFormComponent} from '@schkola/finance/fee/components/form/fee-form.component';
import {FeeKindViewComponent} from '@schkola/finance/feekind/components/view/feekind-entity-view.component';
import {FeeKindListComponent} from '@schkola/finance/feekind/components/list/feekind-entity-list.component';
import {FeeKindFormComponent} from '@schkola/finance/feekind/components/form/feekind-form.component';
import {PersonModule} from '@schkola/person/person-model.module';


@NgModule({
    declarations: [
        FinanceViewComponent,
        ExpenseViewComponent,
        ExpenseListComponent,
        ExpenseFormComponent,
        ExpensePurposeViewComponent,
        ExpensePurposeListComponent,
        ExpensePurposeFormComponent,
        FeeViewComponent,
        FeeListComponent,
        FeeFormComponent,
        FeeKindViewComponent,
        FeeKindListComponent,
        FeeKindFormComponent,

    ],
    imports: [
        FinanceRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule,
        PersonModule
    ],
    providers: [],
    exports: [
        ExpenseFormComponent,
        ExpensePurposeFormComponent,
        FeeFormComponent,
        FeeKindFormComponent,

    ]
})
export class FinanceModule {}



