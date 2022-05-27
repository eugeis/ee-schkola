import {Component, Input, OnInit} from '@angular/core';
import {TemplateService} from '../../services/template.service';

@Component({
    selector: 'app-template',
    templateUrl: './template.component.html',
    styleUrls: ['./template.component.scss'],
    providers: [TemplateService],
})

export class TemplateComponent implements OnInit {

    @Input() formElement: any[][];

    constructor(public templateService: TemplateService) { }

    ngOnInit(): void {
        this.templateService.initElement(this.formElement);
    }
}
