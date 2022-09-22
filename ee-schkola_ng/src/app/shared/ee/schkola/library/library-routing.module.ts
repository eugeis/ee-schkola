import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import {LibraryViewComponent} from './components/view/library-module-view.component';
import {BookViewComponent} from './book/components/view/book-entity-view.component';
import {BookListComponent} from './book/components/list/book-entity-list.component';

const routes: Routes = [
    { path: '', component: LibraryViewComponent },
    { path: 'book', component: BookListComponent },
    { path: 'book/new', component: BookViewComponent },
    { path: 'book/edit/:id', component: BookViewComponent }
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule],
})
export class LibraryRoutingModules {}





