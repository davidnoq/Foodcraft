import { ComponentFixture } from '@angular/core/testing';
import { TestBed, async } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from 'app/auth.service';
import { MatDialogModule } from '@angular/material/dialog';
import { PorkComponent } from './pork.component';


  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, RouterTestingModule, MatDialogModule],
      declarations: [PorkComponent],
      providers: [AuthService],
    }).compileComponents();
  });

  

    describe('PorkComponent', () => {
  it('mounts', () => {
    // change the viewport of the test
    cy.viewport(1300, 700);
    // mount the component to test
   cy.mount(PorkComponent);
    // check visual text
    cy.contains('Generated Recipes');
    

    
  });
  
});
