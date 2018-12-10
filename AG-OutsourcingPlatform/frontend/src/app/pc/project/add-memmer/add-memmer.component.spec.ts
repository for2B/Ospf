import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { AddMemmerComponent } from './add-memmer.component';

describe('AddMemmerComponent', () => {
  let component: AddMemmerComponent;
  let fixture: ComponentFixture<AddMemmerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ AddMemmerComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(AddMemmerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
