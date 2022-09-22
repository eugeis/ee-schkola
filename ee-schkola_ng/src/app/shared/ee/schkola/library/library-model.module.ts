import { NgModule } from '@angular/core';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';

import {LibraryRoutingModules} from './library-routing.module';
import {CommonModule} from '@angular/common';
import {TemplateModule} from '../../template/template.module';
import {MaterialModule} from '../../template/material.module';

import {LibraryViewComponent} from './components/view/library-module-view.component';
import {BookViewComponent} from './book/components/view/book-entity-view.component';
import {BookListComponent} from './book/components/list/book-entity-list.component';

import {LocationComponent} from './basics/location/location-basic.component';
import {PersonModule} from '@schkola/person/person-model.module';

@NgModule({
    declarations: [
        LibraryViewComponent,
        BookViewComponent,
        BookListComponent,
        LocationComponent
    ],
    imports: [
        LibraryRoutingModules,
        TemplateModule,
        CommonModule,
        FormsModule,
        ReactiveFormsModule,
        MaterialModule,
        PersonModule
    ],
    providers: [],
    exports: [
        BookViewComponent,
        LocationComponent
    ]
})
export class LibraryModule { }




