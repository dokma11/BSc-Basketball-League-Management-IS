import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DeclareForDraftComponent } from './declare-for-draft.component';

describe('DeclareForDraftComponent', () => {
  let component: DeclareForDraftComponent;
  let fixture: ComponentFixture<DeclareForDraftComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DeclareForDraftComponent]
    });
    fixture = TestBed.createComponent(DeclareForDraftComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
