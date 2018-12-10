import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { MypublishProjectComponent } from './mypublish-project.component';

describe('MypublishProjectComponent', () => {
  let component: MypublishProjectComponent;
  let fixture: ComponentFixture<MypublishProjectComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ MypublishProjectComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(MypublishProjectComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
